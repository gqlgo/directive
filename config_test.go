package directive

import (
	"github.com/google/go-cmp/cmp"
	"github.com/vektah/gqlparser/v2/ast"
	"os"
	"testing"
)

func TestParseConfigFile(t *testing.T) {
	type args struct {
		configFile string
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "TestParseConfigFile",
			args: args{
				configFile: "./testdata/iddirective/iddirective.yaml",
			},
			want: &Config{
				InputObjectFieldConfig: []*InputObjectFieldConfig{
					{
						Description:  "inputのフィールドにid directiveが存在することをチェックする",
						Directive:    "id",
						Kind:         ast.InputObject,
						FieldType:    "ID",
						ReportFormat: "%s.%s has no id directive",
					},
				},
				ObjectFieldArgumentConfig: []*ObjectFieldArgumentConfig{
					{
						Description:       "fieldの引数にid directiveが存在することをチェックする",
						Directive:         "id",
						Kind:              ast.Object,
						FieldArgumentType: "ID",
						ReportFormat:      "argument %s of %s has no id directive",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configFile, err := os.Open(tt.args.configFile)
			if err != nil {
				t.Fatalf("failed to open config file: %v", err)
			}
			got, err := ParseConfigFile(configFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseConfigFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ParseConfigFile()\n%v", diff)
			}
		})
	}
}
