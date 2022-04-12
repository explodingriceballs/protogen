package modules

import (
	"github.com/dop251/goja"
	"log"
)

type Console struct {
}

func (c *Console) Register(runtime *goja.Runtime) error {
	o := runtime.NewObject()
	if err := runtime.Set("console", o); err != nil {
		return err
	}
	if err := o.Set("log", func(s string) {
		log.Println(s)
	}); err != nil {
		return err
	}

	return nil
}
