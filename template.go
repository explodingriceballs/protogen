package main

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/dop251/goja"
	"log"
	"os"
	"path"
	"protogen/vm"
	"text/template"
)

type TemplateEngine struct {
	templateDir string
}

func (t *TemplateEngine) Create(runtime *vm.Runtime) (*goja.Object, error) {
	object := runtime.CreateObject()
	if err := object.Set("render", t.Render); err != nil {
		return nil, err
	}

	return object, nil
}

func (t *TemplateEngine) Render(fileName string, vars map[string]interface{}) string {
	fullPath := path.Join(t.templateDir, fileName)
	templateBytes, err := os.ReadFile(fullPath)
	if err != nil {
		log.Panic("failed to open file", err)
	}
	tmpl, err := template.New(fileName).Funcs(sprig.TxtFuncMap()).Parse(string(templateBytes))
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

func NewTemplateEngine(templateDir string) *TemplateEngine {
	t := TemplateEngine{
		templateDir: templateDir,
	}
	return &t
}
