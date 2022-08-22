package js

import (
	_ "embed"
	"github.com/dop251/goja"
	"github.com/dop251/goja/parser"
	"sync"
)

//go:embed lib/babel.min.js
var babelSrc string

//go:embed lib/v4.2.4.min.js
var tscSrc string

type Compiler struct {
	babelVm       *goja.Runtime
	babelCompiler goja.Value
	TransformFunc goja.Callable

	tscVm        *goja.Runtime
	tscCompiler  goja.Value
	tscTranspile goja.Callable
	tscModule    goja.Value

	lockMutex sync.Mutex
}

func (c *Compiler) Compile(source string, filename string, main bool) (*goja.Program, error) {
	// Compiling is locked
	c.lockMutex.Lock()
	defer c.lockMutex.Unlock()

	options := map[string]interface{}{
		"compilerOptions": map[string]interface{}{
			"module": c.tscModule,
			//"module": "es6",
			//"target": "es5",
		},
	}

	// TSC test
	value, err := c.tscTranspile(c.tscCompiler, c.tscVm.ToValue(c.tscVm.ToValue(source)), c.tscVm.ToValue(options))
	if err != nil {
		return nil, err
	}
	transpiledSource := value.ToObject(c.tscVm).Get("outputText").String()

	// Transform the code
	opts := c.getOptions()
	opts["filename"] = filename
	v, err := c.TransformFunc(c.babelCompiler, c.babelVm.ToValue(transpiledSource), c.babelVm.ToValue(opts))
	if err != nil {
		return nil, err
	}
	vObject := v.ToObject(c.babelVm)

	var code string
	if err := c.babelVm.ExportTo(vObject.Get("code"), &code); err != nil {
		return nil, err
	}

	if !main {
		code = "(function(module, exports){\n" + code + "\n})\n"
	}

	ast, err := parser.ParseFile(nil, filename, code, 0, parser.WithDisableSourceMaps)
	if err != nil {
		return nil, err
	}

	program, err := goja.CompileAST(ast, true)
	if err != nil {
		return nil, err
	}

	return program, nil
}

func (c *Compiler) getOptions() map[string]interface{} {
	return map[string]interface{}{
		// "presets": []string{"latest"},
		"plugins": []interface{}{
			// es2015 https://github.com/babel/babel/blob/v6.26.0/packages/babel-preset-es2015/src/index.js
			// in goja
			// []interface{}{"transform-es2015-template-literals", map[string]interface{}{"loose": false, "spec": false}},
			// "transform-es2015-literals", // in goja
			// "transform-es2015-function-name", // in goja
			// []interface{}{"transform-es2015-arrow-functions", map[string]interface{}{"spec": false}}, // in goja
			// "transform-es2015-block-scoped-functions", // in goja
			[]interface{}{"transform-es2015-classes", map[string]interface{}{"loose": false}},
			"transform-es2015-object-super",
			// "transform-es2015-shorthand-properties", // in goja
			// "transform-es2015-duplicate-keys", // in goja
			// []interface{}{"transform-es2015-computed-properties", map[string]interface{}{"loose": false}}, // in goja
			// "transform-es2015-for-of", // in goja
			// "transform-es2015-sticky-regex", // in goja
			// "transform-es2015-unicode-regex", // in goja
			// "check-es2015-constants", // in goja
			// []interface{}{"transform-es2015-spread", map[string]interface{}{"loose": false}}, // in goja
			// "transform-es2015-parameters", // in goja
			// []interface{}{"transform-es2015-destructuring", map[string]interface{}{"loose": false}}, // in goja
			// "transform-es2015-block-scoping", // in goja
			// "transform-es2015-typeof-symbol", // in goja
			// all the other module plugins are just dropped
			[]interface{}{"transform-es2015-modules-commonjs", map[string]interface{}{"loose": false}},
			// "transform-regenerator", // Doesn't really work unless regeneratorRuntime is also added

			// es2016 https://github.com/babel/babel/blob/v6.26.0/packages/babel-preset-es2016/src/index.js
			"transform-exponentiation-operator",

			// es2017 https://github.com/babel/babel/blob/v6.26.0/packages/babel-preset-es2017/src/index.js
			// "syntax-trailing-function-commas", // in goja
			// "transform-async-to-generator", // Doesn't really work unless regeneratorRuntime is also added
		},
		"ast":           false,
		"sourceMaps":    false,
		"babelrc":       false,
		"compact":       false,
		"retainLines":   true,
		"highlightCode": false,
	}
}

func (c *Compiler) compileAndLoad(name string, source string, strict bool) (*goja.Runtime, error) {
	program, err := goja.Compile(name, source, strict)
	if err != nil {
		return nil, err
	}
	vm := goja.New()
	_, err = vm.RunProgram(program)
	if err != nil {
		return nil, err
	}

	return vm, nil
}

func (c *Compiler) loadBabel() error {
	babelVm, err := c.compileAndLoad("babel.min.js", babelSrc, false)
	if err != nil {
		return err
	}

	babelCompiler := babelVm.Get("Babel")
	babelObject := babelCompiler.ToObject(babelVm)
	c.babelVm = babelVm
	c.babelCompiler = babelCompiler
	if err := babelVm.ExportTo(babelObject.Get("transform"), &c.TransformFunc); err != nil {
		return err
	}
	return nil
}

func (c *Compiler) loadTsc() error {
	vm, err := c.compileAndLoad("v4.2.4.min.js", tscSrc, false)
	if err != nil {
		return err
	}

	tsc := vm.Get("ts")
	tscObject := tsc.ToObject(vm)
	object := tscObject.Get("ModuleKind").ToObject(c.tscVm).Get("CommonJS").ToNumber()
	c.tscModule = object
	c.tscVm = vm
	c.tscCompiler = tscObject
	if err := vm.ExportTo(tscObject.Get("transpileModule"), &c.tscTranspile); err != nil {
		return err
	}
	return nil
}

func NewCompiler() (*Compiler, error) {
	compiler := Compiler{}
	if err := compiler.loadBabel(); err != nil {
		return nil, err
	}
	if err := compiler.loadTsc(); err != nil {
		return nil, err
	}

	return &compiler, nil
}
