package directive_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gqlgo/directive"
	"github.com/gqlgo/gqlanalysis/analysistest"
)

func TestDirective(t *testing.T) {
	type args struct {
		analyzer string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "constraint directive",
			args: args{analyzer: "constraint"},
		},
		{
			name: "id directive",
			args: args{analyzer: "id"},
		},
		{
			name: "list directive",
			args: args{analyzer: "list"},
		},
		{
			name: "permission directive",
			args: args{analyzer: "permission"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testdata := analysistest.TestData(t)
			configFilePath := fmt.Sprintf("./testdata/%s/directive.yaml", tt.args.analyzer)
			configFile, err := os.Open(configFilePath)
			if err != nil {
				t.Fatalf("failed to open config file: %v", err)
			}
			config, err := directive.ParseConfigFile(configFile)
			if err != nil {
				t.Fatalf("failed to parse config file: %v", err)
			}
			analysistest.Run(t, testdata, directive.NewAnalyzers(config)[0], tt.args.analyzer)
		})
	}
}
