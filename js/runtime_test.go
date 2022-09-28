package js

import (
	"github.com/dop251/goja"
	"gotest.tools/v3/assert"
	"os"
	"testing"
)

type (
	ModuleClass         struct{}
	ModuleClassInstance struct {
		rt *goja.Runtime
		s  string
	}
)

func (m *ModuleClassInstance) Exports() Exports {
	return Exports{
		Named: map[string]interface{}{
			"ModuleClass": m.NewModuleClass,
		},
	}
}

func (m *ModuleClassInstance) NewModuleClass(call goja.ConstructorCall) *goja.Object {
	m.s = call.Argument(0).String()
	call.This.Set("GetString", func() string {
		return "test - " + m.s
	})
	return nil
}

func (m *ModuleClass) Register(runtime *goja.Runtime) error {
	return nil
}

func (m *ModuleClass) NewModuleInstance(runtime *goja.Runtime) Instance {
	return &ModuleClassInstance{rt: runtime}
}

func TestRuntime_RegisterModule(t *testing.T) {
	testFileBytes, err := os.ReadFile("testdata/modules.tsx")
	assert.NilError(t, err)
	runtime := NewRuntime(string(testFileBytes), "testdata/modules.tsx")
	err = runtime.RegisterModule("modules_export", &ModuleClass{})
	assert.NilError(t, err)

	err = runtime.Execute()
	assert.NilError(t, err)

	defaultExport, err := runtime.GetDefaultExport()
	assert.NilError(t, err)

	err = defaultExport.Instantiate()
	assert.NilError(t, err)

	var result string
	err = defaultExport.Invoke("TestFn", &result)
	assert.NilError(t, err)

	assert.Equal(t, result, "test - foobar")
}

func TestRuntime_Require(t *testing.T) {
	type args struct {
		sourceFile string
		testMethod string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    int
	}{
		{
			name: "test basic require",
			args: args{
				sourceFile: "testdata/require.tsx",
				testMethod: "TestMethod",
			},
			wantErr: false,
			want:    10,
		},
		{
			name: "test method require",
			args: args{
				sourceFile: "testdata/require.tsx",
				testMethod: "TestMethod2",
			},
			wantErr: false,
			want:    25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileContents, err := os.ReadFile(tt.args.sourceFile)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			r := NewRuntime(string(fileContents), tt.args.sourceFile)
			if err := r.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err = r.Compile(); err != nil {
				t.Errorf("Compile() error = %v", err)
			}
			if err = r.Execute(); err != nil {
				t.Errorf("Execute() error = %v", err)
			}
			export, err := r.GetDefaultExport()
			if err != nil {
				t.Errorf("GetDefaultExport() error = %v", err)
			}

			if err := export.Instantiate(); err != nil {
				t.Errorf("Instantiate() error = %v", err)
			}

			intValue := 0
			if err := export.Invoke(tt.args.testMethod, &intValue); err != nil {
				t.Errorf("Invoke() error = %v", err)
			}
			if intValue != tt.want {
				t.Errorf("Invoke() retValue = %v, want %v", intValue, tt.want)
			}
		})
	}
}
