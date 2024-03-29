package js

import (
	"github.com/dop251/goja"
	"github.com/explodingriceballs/protogen/js/modules"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Compiler Pool
var compilerPoolLogger = zap.L()
var compilerPool = sync.Pool{
	New: func() interface{} {
		compiler, err := NewCompiler()
		if err != nil && compilerPoolLogger != nil {
			compilerPoolLogger.Info("compiler pool error", zap.Error(err))
		}
		return compiler
	},
}

var fileExt = []string{".d.ts", ".tsx", ".ts"}

type Runtime struct {
	source     string
	sourceFile string
	compiled   bool
	vm         *goja.Runtime
	program    *goja.Program
	exports    *goja.Object
	modules    map[string]Module
}

func (r *Runtime) Compile() error {
	// Don't compile twice
	if r.compiled {
		return nil
	}

	// Get a new compiler from the pool
	compiler := compilerPool.Get().(*Compiler)
	defer func() {
		// Return the compiler to the pool
		compilerPool.Put(compiler)
	}()

	program, err := compiler.Compile(r.source, r.sourceFile, true)
	if err != nil {
		return err
	}

	// Save the compiled outputs
	r.program = program

	// Set to compiled
	r.compiled = true

	return nil
}

func (r *Runtime) Execute() error {
	if !r.compiled {
		if err := r.Compile(); err != nil {
			return err
		}
	}

	console := modules.Console{}
	if err := console.Register(r.vm); err != nil {
		return err
	}

	// Dummy "require" func
	if err := r.vm.Set("require", func(arg string) goja.Value {
		v, err := r.requireFile(arg)
		if err != nil {
			r.PanicError(err.Error())
		}
		return v
	}); err != nil {
		return err
	}

	if r.vm.Get("exports") == nil {
		r.exports = r.vm.NewObject()
		if err := r.vm.Set("exports", r.exports); err != nil {
			return err
		}
	}

	_, err := r.vm.RunProgram(r.program)
	return err
}

func (r *Runtime) RegisterModule(s string, m Module) error {
	err := m.Register(r.vm)
	r.modules[s] = m
	return err
}

func (r *Runtime) CreateObject() *goja.Object {
	return r.vm.NewObject()
}

func (r *Runtime) GetDefaultExport() (*Class, error) {
	var (
		constructor goja.Callable
	)

	defaultObj := r.exports.Get("default").ToObject(r.vm)
	if err := r.vm.ExportTo(defaultObj, &constructor); err != nil {
		return nil, err
	}

	prototype := defaultObj.Get("prototype").ToObject(r.vm)
	class := NewClass(r.vm, constructor, prototype)
	return class, nil
}

func (r *Runtime) PanicError(error string) {
	panic(r.vm.ToValue(error))
}

func (r *Runtime) Close() error {
	r.vm.Interrupt("halt")
	return nil
}

func (r *Runtime) GetRuntime() *goja.Runtime {
	return r.vm
}

func (r *Runtime) requireFile(name string) (goja.Value, error) {
	// Base directory of the original source file (absolute)
	baseDir, err := filepath.Abs(filepath.Dir(r.sourceFile))
	if err != nil {
		return nil, err
	}

	// Join the base dir & the required file
	file := filepath.Join(baseDir, name)

	// Module name
	moduleName := strings.Replace(file, baseDir, "", 1)
	if moduleName[0] == '/' {
		moduleName = moduleName[1:]
	}
	if v, ok := r.modules[moduleName]; ok {
		instance := v.NewModuleInstance(r.vm)
		return r.vm.ToValue(toESModuleExports(instance.Exports())), nil
	}

	if filepath.Ext(name) == "" {
		for _, ext := range fileExt {
			if _, err := os.Stat(file + ext); os.IsNotExist(err) {
				continue
			}
			file = file + ext
			break
		}
	}

	fileSource, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// Get a new compiler from the pool
	compiler := compilerPool.Get().(*Compiler)
	defer func() {
		// Return the compiler to the pool
		compilerPool.Put(compiler)
	}()

	program, err := compiler.Compile(string(fileSource), file, false)
	if err != nil {
		return nil, err
	}

	f, err := r.vm.RunProgram(program)
	if err != nil {
		return nil, err
	}

	exports := r.vm.NewObject()
	module := r.vm.NewObject()
	_ = module.Set("exports", exports)
	if call, ok := goja.AssertFunction(f); ok {
		if _, err := call(exports, module, exports); err != nil {
			return nil, err
		}
	}

	return exports, nil
}

// toESModuleExports & the export system for including native module is largely based on the code from k6:
// https://github.com/grafana/k6/blob/d19102983eeaedf9d53e78ba73bb17f2357fc522/js/initcontext.go#L186
func toESModuleExports(exp Exports) interface{} {
	if exp.Named == nil {
		return exp.Default
	}
	if exp.Default == nil {
		return exp.Named
	}

	result := make(map[string]interface{}, len(exp.Named)+2)

	for k, v := range exp.Named {
		result[k] = v
	}
	// Maybe check that those weren't set
	result["default"] = exp.Default
	// this so babel works with the `default` when it transpiles from ESM to commonjs.
	// This should probably be removed once we have support for ESM directly. So that require doesn't get support for
	// that while ESM has.
	result["__esModule"] = true

	return result
}

func NewRuntime(source string, sourceFile string) *Runtime {
	return &Runtime{
		source:     source,
		sourceFile: sourceFile,
		compiled:   false,
		vm:         goja.New(),
		modules:    map[string]Module{},
	}
}
