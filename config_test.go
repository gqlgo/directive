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
			name: "constraint directive",
			args: args{
				configFile: "./testdata/constraint/directive.yaml",
			},
			want: &Config{
				Analyzer: []*AnalyzerConfig{
					{
						AnalyzerName: "constraint directive",
						Description:  "constraint directive exists on the field",
						InputObjectFieldConfig: []*InputObjectFieldConfig{
							{
								Description:             "constraint directive exists on the input field",
								Directive:               "constraint",
								FieldTypePatterns:       []string{`^\[?Int\]?$`, `^\[?Float\]?$`, `^\[?String\]?$`, `^\[?Decimal\]?$`, `^\[?URL\]?$`},
								IgnoreFieldNamePatterns: []string{`^first$`, `^last$`, `^after$`, `^before$`},
								ReportFormat:            "%s.%s has no constraint directive",
							},
						},
						ObjectFieldArgumentConfig: []*ObjectFieldArgumentConfig{
							{
								Description:                "constraint directive exists on the object field argument",
								Directive:                  "constraint",
								ArgumentTypePatterns:       []string{`^\[?Int\]?$`, `^\[?Float\]?$`, `^\[?String\]?$`, `^\[?Decimal\]?$`, `^\[?URL\]?$`},
								IgnoreArgumentNamePatterns: []string{`^first$`, `^last$`, `^after$`, `^before$`},
								ReportFormat:               "argument %s of %s has no constraint directive",
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "id directive",
			args: args{
				configFile: "./testdata/id/directive.yaml",
			},
			want: &Config{
				Analyzer: []*AnalyzerConfig{
					{
						AnalyzerName: "id directive",
						Description:  "id directive exists on the field",
						InputObjectFieldConfig: []*InputObjectFieldConfig{
							{
								Description:       "id directive exists on the input field",
								Directive:         "id",
								FieldTypePatterns: []string{`^\[?ID\]?$`},
								ReportFormat:      "%s.%s has no id directive",
							},
						},
						ObjectFieldArgumentConfig: []*ObjectFieldArgumentConfig{
							{
								Description:          "id directive exists on the object field argument",
								Directive:            "id",
								ArgumentTypePatterns: []string{`^\[?ID\]?$`},
								ReportFormat:         "argument %s of %s has no id directive",
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "list directive",
			args: args{
				configFile: "./testdata/list/directive.yaml",
			},
			want: &Config{
				Analyzer: []*AnalyzerConfig{
					{
						AnalyzerName: "list directive",
						Description:  "list directive exists on the array field",
						InputObjectFieldConfig: []*InputObjectFieldConfig{
							{
								Description:       "list directive exists on the input array field",
								Directive:         "list",
								FieldTypePatterns: []string{`\[.+\]`},
								ReportFormat:      "%s.%s has no list directive",
							},
						},
						ObjectFieldArgumentConfig: []*ObjectFieldArgumentConfig{
							{
								Description:          "list directive exists on the object field array argument",
								Directive:            "list",
								ArgumentTypePatterns: []string{`\[.+\]`},
								ReportFormat:         "argument %s of %s has no list directive",
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "permission directive",
			args: args{
				configFile: "./testdata/permission/directive.yaml",
			},
			want: &Config{
				Analyzer: []*AnalyzerConfig{
					{
						AnalyzerName: "permission directive",
						Description:  "permission directive exists on the type",
						TypeConfig: []*TypeConfig{
							{
								Description:        "permission directive exists on the type",
								Directive:          "permission",
								Kinds:              []ast.DefinitionKind{"OBJECT", "INTERFACE", "FIELD_DEFINITION"},
								ObjectPatterns:     []string{".*"},
								IgnoreTypePatterns: []string{"^Query$", "^Mutation$", "^Subscription$", "PageInfo$", "Connection$", "Payload$"},
								ReportFormat:       "%s has no permission directive",
							},
						},
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
