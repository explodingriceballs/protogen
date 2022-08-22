package vm

import "github.com/dop251/goja"

type Module interface {
	Register(runtime *goja.Runtime) error
	NewModuleInstance(runtime *goja.Runtime) Instance
}

type Instance interface {
	Exports() Exports
}

type Exports struct {
	Default interface{}
	Named   map[string]interface{}
}
