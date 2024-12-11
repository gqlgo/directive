package directive

import (
	"fmt"
	"github.com/goccy/go-yaml"
	"github.com/vektah/gqlparser/v2/ast"
	"io"
)

type Config struct {
	Analyzer []*AnalyzerConfig `yaml:"analyzer"`
}

type AnalyzerConfig struct {
	AnalyzerName              string                       `yaml:"analyzer_name"`
	Description               string                       `yaml:"description"`
	InputObjectFieldConfig    []*InputObjectFieldConfig    `yaml:"input_object_field"`
	ObjectFieldArgumentConfig []*ObjectFieldArgumentConfig `yaml:"object_field_argument"`
	TypeConfig                []*TypeConfig                `yaml:"type"`
}

type InputObjectFieldConfig struct {
	Description string `yaml:"description"`
	Directive   string `yaml:"directive"`
	// FieldTypePatterns is the name of the field to check.
	FieldTypePatterns       []string `yaml:"field_type"`
	IgnoreFieldNamePatterns []string `yaml:"ignore_field_name"`
	ReportFormat            string   `yaml:"report_format"`
}

type ObjectFieldArgumentConfig struct {
	Description string `yaml:"description"`
	Directive   string `yaml:"directive"`
	// ArgumentTypePatterns is the name of the field argument to check.
	ArgumentTypePatterns       []string `yaml:"argument_type"`
	IgnoreArgumentNamePatterns []string `yaml:"ignore_argument_name"`
	ReportFormat               string   `yaml:"report_format"`
}

type TypeConfig struct {
	Description        string               `yaml:"description"`
	Directive          string               `yaml:"directive"`
	Kinds              []ast.DefinitionKind `yaml:"kind"`
	ObjectPatterns     []string             `yaml:"object_patterns"`
	IgnoreTypePatterns []string             `yaml:"ignore_type"`
	ReportFormat       string               `yaml:"report_format"`
}

func ParseConfigFile(configFile io.Reader) (*Config, error) {
	var config Config
	if err := yaml.NewDecoder(configFile).Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}
	return &config, nil
}
