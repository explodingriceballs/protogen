package vm

import (
	"os"
	"testing"
)

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
