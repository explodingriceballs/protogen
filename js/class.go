package js

import (
	"errors"
	"github.com/dop251/goja"
)

type Class struct {
	vm          *goja.Runtime
	constructor goja.Callable
	prototype   *goja.Object
	this        *goja.Object
	callables   map[string]goja.Callable
}

func NewClass(vm *goja.Runtime, constructor goja.Callable, prototype *goja.Object) *Class {
	return &Class{
		vm:          vm,
		constructor: constructor,
		prototype:   prototype,
		callables:   map[string]goja.Callable{},
	}
}

// Instantiate calls the constructor with the supplied arguments
func (c *Class) Instantiate(params ...interface{}) error {
	// Create a new this
	c.this = c.vm.NewObject()

	// Set the prototype (otherwise internal calls don't work)
	if err := c.this.SetPrototype(c.prototype); err != nil {
		return err
	}
	_, err := c.constructor(c.this, c.convertParams(params...)...)

	// Create a list of callable functions
	for _, k := range c.prototype.Keys() {
		fn, b := goja.AssertFunction(c.prototype.Get(k))
		if !b {
			continue
		}
		c.callables[k] = fn
	}

	return err
}

func (c *Class) InvokeVoid(fnName string, params ...interface{}) error {
	fn, ok := c.callables[fnName]
	if !ok {
		return errors.New("no such function: " + fnName)
	}
	_, err := fn(c.this, c.convertParams(params...)...)
	return err
}

func (c *Class) Invoke(fnName string, result interface{}, params ...interface{}) error {
	fn, ok := c.callables[fnName]
	if !ok {
		return errors.New("no such function: " + fnName)
	}
	ret, err := fn(c.this, c.convertParams(params...)...)
	if err != nil {
		return err
	}
	if err := c.vm.ExportTo(ret, result); err != nil {
		return err
	}

	return err
}

func (c *Class) This() *goja.Object {
	return c.this
}

func (c *Class) convertParams(params ...interface{}) []goja.Value {
	vs := make([]goja.Value, len(params))
	for idx, param := range params {
		vs[idx] = c.vm.ToValue(param)
	}
	return vs
}
