package main

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/dop251/goja"
	"github.com/explodingriceballs/protogen/vm"
	"log"
	"os"
	"path"
	"text/template"
)

type TemplateEngine struct {
	templateDir string
	runtime     *vm.Runtime
	script      *vm.Class
}

func (t *TemplateEngine) Create(runtime *vm.Runtime) (*goja.Object, error) {
	t.runtime = runtime
	object := runtime.CreateObject()
	if err := object.Set("render", t.Render); err != nil {
		return nil, err
	}

	return object, nil
}

func (t *TemplateEngine) Render(fileName string, vars map[string]interface{}, functions map[string]goja.Value) string {
	fullPath := path.Join(t.templateDir, fileName)
	templateBytes, err := os.ReadFile(fullPath)
	if err != nil {
		log.Panic("failed to open file", err)
	}

	scriptFunctions := map[string]any{}
	for key, value := range functions {
		//fnValue := t.runtime.GetRuntime().ToValue(value)
		if callable, ok := goja.AssertFunction(value); ok {
			scriptFunctions[key] = func(args ...interface{}) (interface{}, error) {
				vmArgs := make([]goja.Value, len(args))
				for i, arg := range args {
					vmArgs[i] = t.runtime.GetRuntime().ToValue(arg)
				}
				//export, _ := t.runtime.GetDefaultExport()
				return callable(t.script.This(), vmArgs...)
			}
		}
	}

	tmpl, err := template.New(fileName).
		Funcs(sprig.TxtFuncMap()).
		Funcs(scriptFunctions).
		Parse(string(templateBytes))
	if err != nil {
		log.Panic("failed to parse template", err)
	}

	// Render to buffer & return
	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, vars)
	if err != nil {
		log.Panic("failed to render template", err)
	}
	return tpl.String()
}

func NewTemplateEngine(templateDir string, script *vm.Class) *TemplateEngine {
	t := TemplateEngine{
		templateDir: templateDir,
		script:      script,
	}
	return &t
}
