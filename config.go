package directive

import (
	"fmt"
	"github.com/goccy/go-yaml"
	"github.com/vektah/gqlparser/v2/ast"
	"io"
)

type Config struct {
	InputObjectFieldConfig    []*InputObjectFieldConfig    `yaml:"input_object_field"`
	ObjectFieldArgumentConfig []*ObjectFieldArgumentConfig `yaml:"object_field_argument"`
}

type InputObjectFieldConfig struct {
	Description string `yaml:"description"`
	Directive   string `yaml:"directive"`
	// Kinds is a list of definition kinds to check.
	Kind ast.DefinitionKind `yaml:"kind"`
	// FieldType is the name of the field to check.
	FieldType    string `yaml:"field_type"`
	ReportFormat string `yaml:"report_format"`
}

type ObjectFieldArgumentConfig struct {
	Description string `yaml:"description"`
	Directive   string `yaml:"directive"`
	// Kinds is a list of definition kinds to check.
	Kind ast.DefinitionKind `yaml:"kind"`
	// FieldArgumentType is the name of the field argument to check.
	FieldArgumentType string `yaml:"field_argument_type"`
	ReportFormat      string `yaml:"report_format"`
}

func ParseConfigFile(configFile io.Reader) (*Config, error) {
	var config Config
	if err := yaml.NewDecoder(configFile).Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}
	return &config, nil
}
