package js

import "github.com/dop251/goja"

func Panic(rt *goja.Runtime, err error) {
	value := rt.ToValue(err)
	panic(value)
}
