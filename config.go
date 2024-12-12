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
	AnalyzerName   string            `yaml:"analyzer_name"`
	Description    string            `yaml:"description"`
	FieldConfig    []*FieldConfig    `yaml:"field"`
	ArgumentConfig []*ArgumentConfig `yaml:"argument"`
	TypeConfig     []*TypeConfig     `yaml:"type"`
}

type FieldConfig struct {
	Description             string               `yaml:"description"`
	Directive               string               `yaml:"directive"`
	Kinds                   []ast.DefinitionKind `yaml:"kind"`
	FieldParentTypePatterns []string             `yaml:"field_parent_type"`
	FieldTypePatterns       []string             `yaml:"field_type"`
	IgnoreFieldPatterns     []string             `yaml:"ignore_field"`
	ReportFormat            string               `yaml:"report_format"`
}

type ArgumentConfig struct {
	Description            string               `yaml:"description"`
	Directive              string               `yaml:"directive"`
	Kinds                  []ast.DefinitionKind `yaml:"kind"`
	ArgumentTypePatterns   []string             `yaml:"argument_type"`
	IgnoreArgumentPatterns []string             `yaml:"ignore_argument"`
	ReportFormat           string               `yaml:"report_format"`
}

type TypeConfig struct {
	Description        string               `yaml:"description"`
	Directive          string               `yaml:"directive"`
	Kinds              []ast.DefinitionKind `yaml:"kind"`
	TypePatterns       []string             `yaml:"type"`
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
