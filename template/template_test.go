package template

import (
	"github.com/dop251/goja"
	"github.com/explodingriceballs/protogen/js"
	"gotest.tools/v3/assert"
	"testing"
)

func TestNewTemplateEngine(t *testing.T) {
	const SCRIPT = `
import {Template} from "protogen/template";
function test() {
	let t = new Template("service.go.tmpl");
	return t.TemplateName();
}
`
	runtime := js.NewRuntime(SCRIPT, "")
	err := runtime.RegisterModule("protogen/template", &Module{})
	assert.NilError(t, err)

	err = runtime.Execute()
	assert.NilError(t, err)

	testFn, ok := goja.AssertFunction(runtime.GetRuntime().Get("test"))
	assert.Assert(t, ok == true)

	fnResult, err := testFn(nil)
	assert.NilError(t, err)
	assert.Equal(t, fnResult.ToBoolean(), true)
}

func TestTemplateEngine_ExecuteSimpleVar(t *testing.T) {

	const SCRIPT = `
import {Template} from "protogen/template";
function test(): string {
	let t = new Template("simple_var.tmpl");
	t.Set("var1", "test");
	t.Set("arrayVar", ["1", "2"]);
	return t.Execute();
}
`
	runtime := js.NewRuntime(SCRIPT, "")
	err := runtime.RegisterModule("protogen/template", &Module{TemplateDir: "testdata"})
	assert.NilError(t, err)

	err = runtime.Execute()
	assert.NilError(t, err)

	testFn, ok := goja.AssertFunction(runtime.GetRuntime().Get("test"))
	assert.Assert(t, ok == true)

	fnResult, err := testFn(nil)
	assert.NilError(t, err)
	assert.Equal(t, fnResult.ToString().String(), `Var1: test
Array: 1,2,`)
}

func TestTemplateEngine_ExecuteSimpleFunc(t *testing.T) {

	const SCRIPT = `
import {Template} from "protogen/template";
function test() {
	let t = new Template("simple_func.tmpl");
	t.Set("var1", "foobar");
	t.RegisterFunc("helloWorld", function() {
		return "Hello World!";
	});
	t.RegisterFunc("testFunc", function(s: string) {
		return "bar-" + s;
	});
	return t.Execute();
}
`
	runtime := js.NewRuntime(SCRIPT, "")
	err := runtime.RegisterModule("protogen/template", &Module{TemplateDir: "testdata"})
	assert.NilError(t, err)

	err = runtime.Execute()
	assert.NilError(t, err)

	testFn, ok := goja.AssertFunction(runtime.GetRuntime().Get("test"))
	assert.Assert(t, ok == true)

	fnResult, err := testFn(nil)
	assert.NilError(t, err)
	assert.Equal(t, fnResult.ToString().String(), `Test Func: Hello World!
Param Func: bar-foobar`)
}

func TestTemplateEngine_ExecuteClassTest(t *testing.T) {

	const SCRIPT = `
import {Template} from "protogen/template";
export default class TestClass {
	TestFunction(): string {
		let tmpl = new Template("simple_var.tmpl");
		tmpl.Set("var1", "test");
		tmpl.Set("arrayVar", ["1", "2"]);
		return tmpl.Execute();
	}
}`
	runtime := js.NewRuntime(SCRIPT, "")
	err := runtime.RegisterModule("protogen/template", &Module{TemplateDir: "testdata"})
	assert.NilError(t, err)

	err = runtime.Execute()
	assert.NilError(t, err)

	export, err := runtime.GetDefaultExport()
	assert.NilError(t, err)
	err = export.Instantiate()
	assert.NilError(t, err)

	var fnResult string
	err = export.Invoke("TestFunction", &fnResult)
	assert.NilError(t, err)
	assert.Equal(t, fnResult, `Var1: test
Array: 1,2,`)
}
