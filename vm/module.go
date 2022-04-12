package vm

import "github.com/dop251/goja"

type Module interface {
	Register(runtime *goja.Runtime) error
}

type Object interface {
	Create(runtime *Runtime) (*goja.Object, error)
}
