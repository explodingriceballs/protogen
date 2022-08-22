package template

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/dop251/goja"
	"github.com/explodingriceballs/protogen/js"
	"os"
	"path"
	"text/template"
)

type (
	Module struct {
		TemplateDir string
	}
	ModuleInstance struct {
		rt          *goja.Runtime
		templateDir string
	}
	Template struct {
		rt          *goja.Runtime
		templateDir string
		template    string
		vars        map[string]any
		funcs       template.FuncMap
	}
)

func (m *Module) Register(runtime *goja.Runtime) error {
	return nil
}

func (m *Module) NewModuleInstance(runtime *goja.Runtime) js.Instance {
	return &ModuleInstance{rt: runtime, templateDir: m.TemplateDir}
}

func (m *ModuleInstance) Exports() js.Exports {
	return js.Exports{
		Named: map[string]interface{}{
			"Template": m.NewTemplate,
		},
	}
}

func (m *ModuleInstance) NewTemplate(call goja.ConstructorCall) *goja.Object {
	tmpl := &Template{
		rt:          m.rt,
		templateDir: m.templateDir,
		template:    call.Argument(0).String(),
		vars:        map[string]any{},
		funcs:       map[string]any{},
	}
	return m.rt.ToValue(tmpl).ToObject(m.rt)
}

//func (t *TemplateEngine) Register(runtime *goja.Runtime) error {
//	err := runtime.Set("Template", func(call goja.ConstructorCall) *goja.Object {
//		tmpl := NewTemplate(t.TemplateDir, call.Argument(0).String())
//		_ = call.This.Set("TemplateName", func(_ goja.FunctionCall) goja.Value {
//			return runtime.ToValue(tmpl.TemplateName())
//		})
//		_ = call.This.Set("Execute", func(call goja.FunctionCall) goja.Value {
//			err, compiledTmpl := tmpl.Execute()
//			if err != nil {
//				panic(runtime.ToValue(err))
//			}
//			return runtime.ToValue(compiledTmpl)
//		})
//		_ = call.This.Set("Set", func(name string, val goja.Value) {
//			tmpl.Set(name, val.Export())
//		})
//		_ = call.This.Set("RegisterFunc", func(name string, callable goja.Callable) {
//			tmpl.RegisterFunc(name, func(params ...interface{}) string {
//				nativeVals := make([]goja.Value, len(params))
//				for idx, param := range params {
//					nativeVals[idx] = runtime.ToValue(param)
//				}
//				value, err := callable(nil, nativeVals...)
//				if err != nil {
//					panic(runtime.ToValue(err))
//				}
//				return value.String()
//			})
//		})
//		return nil
//	})
//
//	return err
//}

func (t *Template) TemplateName() string {
	return path.Join(t.templateDir, t.template)
}

func (t *Template) RegisterFunc(name string, fn func(params ...interface{}) string) {
	t.funcs[name] = fn
}

func (t *Template) Set(name string, val any) {
	t.vars[name] = val
}

func (t *Template) Execute() string {
	templateBytes, err := os.ReadFile(t.TemplateName())
	if err != nil {
		js.Panic(t.rt, err)
	}

	// create a new template
	tmpl, err := template.New(t.TemplateName()).
		Funcs(sprig.TxtFuncMap()).
		Funcs(t.funcs).
		Parse(string(templateBytes))
	if err != nil {
		js.Panic(t.rt, err)
	}

	// Render to buffer & return
	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, t.vars)
	if err != nil {
		js.Panic(t.rt, err)
	}
	return tpl.String()
}
