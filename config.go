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
	AnalyzerName     string              `yaml:"analyzer_name"`
	Description      string              `yaml:"description"`
	DefinitionConfig []*DefinitionConfig `yaml:"definition"`
	FieldConfig      []*FieldConfig      `yaml:"field"`
	ArgumentConfig   []*ArgumentConfig   `yaml:"argument"`
}

type DefinitionConfig struct {
	Description               string               `yaml:"description"`
	Directive                 string               `yaml:"directive"`
	Kinds                     []ast.DefinitionKind `yaml:"kind"`
	DefinitionPatterns        []string             `yaml:"definition"`
	ExcludeDefinitionPatterns []string             `yaml:"exclude_definition"`
	ReportFormat              string               `yaml:"report_format"`
}

type FieldConfig struct {
	Description             string               `yaml:"description"`
	Directive               string               `yaml:"directive"`
	Kinds                   []ast.DefinitionKind `yaml:"kind"`
	FieldParentTypePatterns []string             `yaml:"field_parent_type"`
	FieldTypePatterns       []string             `yaml:"field_type"`
	ExcludeFieldPatterns    []string             `yaml:"exclude_field"`
	ReportFormat            string               `yaml:"report_format"`
}

type ArgumentConfig struct {
	Description             string               `yaml:"description"`
	Directive               string               `yaml:"directive"`
	Kinds                   []ast.DefinitionKind `yaml:"kind"`
	ArgumentTypePatterns    []string             `yaml:"argument_type"`
	ExcludeArgumentPatterns []string             `yaml:"exclude_argument"`
	ReportFormat            string               `yaml:"report_format"`
}

func ParseConfigFile(configFile io.Reader) (*Config, error) {
	var config Config
	if err := yaml.NewDecoder(configFile).Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}
	return &config, nil
}
