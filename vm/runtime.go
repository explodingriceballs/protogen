package vm

import (
	"github.com/dop251/goja"
	"github.com/explodingriceballs/protogen/vm/modules"
	"go.uber.org/zap"
	"os"
	"path/filepath"
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

func (r *Runtime) requireFile(name string) (goja.Value, error) {
	// Base directory of the original source file (absolute)
	baseDir, err := filepath.Abs(filepath.Dir(r.sourceFile))
	if err != nil {
		return nil, err
	}

	file := filepath.Join(baseDir, name)
	if filepath.Ext(name) == "" {
		for _, ext := range fileExt {
			if _, err := os.Stat(file + ext); os.IsNotExist(err) {
				continue
			}
			file = file + ext
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

func NewRuntime(source string, sourceFile string) *Runtime {
	return &Runtime{
		source:     source,
		sourceFile: sourceFile,
		compiled:   false,
		vm:         goja.New(),
	}
}
