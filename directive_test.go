package directive_test

import (
	"os"
	"testing"

	"github.com/gqlgo/directive"
	"github.com/gqlgo/gqlanalysis/analysistest"
)

func TestInputObjectField(t *testing.T) {
	testdata := analysistest.TestData(t)
	iddirectiveConfigFile := "./testdata/iddirectiveinputobjectfield/iddirective.yaml"
	configFile, err := os.Open(iddirectiveConfigFile)
	if err != nil {
		t.Fatalf("failed to open config file: %v", err)
	}
	config, err := directive.ParseConfigFile(configFile)
	if err != nil {
		t.Fatalf("failed to parse config file: %v", err)
	}
	analysistest.Run(t, testdata, directive.InputObjectFieldAnalyzer(config.InputObjectFieldConfig[0]), "iddirectiveinputobjectfield")
}

func TestObjectFieldArgument(t *testing.T) {
	testdata := analysistest.TestData(t)
	iddirectiveConfigFile := "./testdata/iddirectiveobjectfieldargument/iddirective.yaml"
	configFile, err := os.Open(iddirectiveConfigFile)
	if err != nil {
		t.Fatalf("failed to open config file: %v", err)
	}
	config, err := directive.ParseConfigFile(configFile)
	if err != nil {
		t.Fatalf("failed to parse config file: %v", err)
	}
	analysistest.Run(t, testdata, directive.ObjectFieldArgumentAnalyzer(config.ObjectFieldArgumentConfig[0]), "iddirectiveobjectfieldargument")
}
