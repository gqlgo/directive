package directive_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gqlgo/directive"
	"github.com/gqlgo/gqlanalysis/analysistest"
)

func TestDirective(t *testing.T) {
	targets := []string{"constraint", "id", "list", "permission"}
	for _, target := range targets {
		testdata := analysistest.TestData(t)
		configFilePath := fmt.Sprintf("./testdata/%s/directive.yaml", target)
		configFile, err := os.Open(configFilePath)
		if err != nil {
			t.Fatalf("failed to open config file: %v", err)
		}
		config, err := directive.ParseConfigFile(configFile)
		if err != nil {
			t.Fatalf("failed to parse config file: %v", err)
		}
		analysistest.Run(t, testdata, directive.NewAnalyzers(config)[0], target)
	}
}
