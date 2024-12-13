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
						FieldConfig: []*FieldConfig{
							{
								Description:             "constraint directive exists on the input field",
								Directive:               "constraint",
								Kinds:                   []ast.DefinitionKind{ast.InputObject},
								FieldParentTypePatterns: []string{`.+`},
								FieldTypePatterns:       []string{`^\[?Int\]?$`, `^\[?Float\]?$`, `^\[?String\]?$`, `^\[?Decimal\]?$`, `^\[?URL\]?$`},
								ExcludeFieldPatterns:    []string{`^first$`, `^last$`, `^after$`, `^before$`},
								ReportFormat:            "%s has no constraint directive",
							},
						},
						ArgumentConfig: []*ArgumentConfig{
							{
								Description:             "constraint directive exists on the object field argument",
								Directive:               "constraint",
								Kinds:                   []ast.DefinitionKind{ast.Object},
								ArgumentTypePatterns:    []string{`^\[?Int\]?$`, `^\[?Float\]?$`, `^\[?String\]?$`, `^\[?Decimal\]?$`, `^\[?URL\]?$`},
								ExcludeArgumentPatterns: []string{`^first$`, `^last$`, `^after$`, `^before$`},
								ReportFormat:            "argument %s has no constraint directive",
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
						FieldConfig: []*FieldConfig{
							{
								Description:             "id directive exists on the input field",
								Directive:               "id",
								Kinds:                   []ast.DefinitionKind{ast.InputObject},
								FieldParentTypePatterns: []string{`.+`},
								FieldTypePatterns:       []string{`^\[?ID\]?$`},
								ReportFormat:            "%s has no id directive",
							},
						},
						ArgumentConfig: []*ArgumentConfig{
							{
								Description:          "id directive exists on the object field argument",
								Directive:            "id",
								Kinds:                []ast.DefinitionKind{ast.Object},
								ArgumentTypePatterns: []string{`^\[?ID\]?$`},
								ReportFormat:         "argument %s has no id directive",
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
						FieldConfig: []*FieldConfig{
							{
								Description:             "list directive exists on the input array field",
								Directive:               "list",
								Kinds:                   []ast.DefinitionKind{ast.InputObject},
								FieldParentTypePatterns: []string{`.+`},
								FieldTypePatterns:       []string{`\[.+\]`},
								ReportFormat:            "%s has no list directive",
							},
						},
						ArgumentConfig: []*ArgumentConfig{
							{
								Description:          "list directive exists on the object field array argument",
								Directive:            "list",
								Kinds:                []ast.DefinitionKind{ast.Object},
								ArgumentTypePatterns: []string{`\[.+\]`},
								ReportFormat:         "argument %s has no list directive",
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
						Description:  "permission directive exists on the definition",
						DefinitionConfig: []*DefinitionConfig{
							{
								Description:               "permission directive exists on the definition",
								Directive:                 "permission",
								Kinds:                     []ast.DefinitionKind{ast.Object, ast.Interface},
								DefinitionPatterns:        []string{".+"},
								ExcludeDefinitionPatterns: []string{"^Query$", "^Mutation$", "^Subscription$", "^PageInfo$"},
								ReportFormat:              "%s has no permission directive",
							},
						},
						FieldConfig: []*FieldConfig{
							{
								Description:             "permission directive exists on the mutation",
								Directive:               "permission",
								Kinds:                   []ast.DefinitionKind{ast.Object},
								FieldParentTypePatterns: []string{`^Mutation$`},
								FieldTypePatterns:       []string{".+"},
								ReportFormat:            "%s has no permission directive",
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
