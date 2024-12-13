package directive

import (
	"github.com/gqlgo/gqlanalysis"
)

func NewAnalyzers(config *Config) []*gqlanalysis.Analyzer {
	analyzers := make([]*gqlanalysis.Analyzer, 0, len(config.Analyzer))
	for _, analyzerConfig := range config.Analyzer {
		analyzers = append(analyzers, NewAnalyzer(analyzerConfig))
	}
	return analyzers
}

func NewAnalyzer(analyzerConfig *AnalyzerConfig) *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: analyzerConfig.AnalyzerName,
		Doc:  analyzerConfig.Description,
		Run:  MergeAnalyzers(analyzerConfig),
	}
}

func MergeAnalyzers(analyzerConfig *AnalyzerConfig) func(pass *gqlanalysis.Pass) (any, error) {
	var analyzers []*gqlanalysis.Analyzer
	for _, c := range analyzerConfig.DefinitionConfig {
		analyzer := DefinitionAnalyzer(c)
		analyzers = append(analyzers, analyzer)
	}
	for _, c := range analyzerConfig.FieldConfig {
		analyzer := FieldAnalyzer(c)
		analyzers = append(analyzers, analyzer)
	}
	for _, c := range analyzerConfig.ArgumentConfig {
		analyzer := ArgumentAnalyzer(c)
		analyzers = append(analyzers, analyzer)
	}
	return func(pass *gqlanalysis.Pass) (any, error) {
		for _, analyzer := range analyzers {
			if _, err := analyzer.Run(pass); err != nil {
				return nil, err
			}
		}
		return nil, nil
	}
}

func DefinitionAnalyzer(config *DefinitionConfig) *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: config.Directive,
		Doc:  config.Description,
		Run: func(pass *gqlanalysis.Pass) (any, error) {
			definitions := NewDefinitionsByMap(pass.Schema.Types).
				NotBuildIn().
				FilterByKinds(config.Kinds).
				ExcludeByDefinitionName(config.ExcludeDefinitionPatterns).
				FilterByNotHasDirective(config.Directive).
				FilterByPositionNotNil()

			for _, definition := range definitions {
				pass.Reportf(definition.Position, config.ReportFormat, definition.Name)
			}

			return nil, nil
		},
	}
}

func FieldAnalyzer(config *FieldConfig) *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: config.Directive,
		Doc:  config.Description,
		Run: func(pass *gqlanalysis.Pass) (any, error) {
			definitions := NewDefinitionsByMap(pass.Schema.Types).
				NotBuildIn().
				FilterByKinds(config.Kinds).
				FilterByDefinitionName(config.FieldParentTypePatterns)
			fields := definitions.Fields().
				FilterByFieldType(config.FieldTypePatterns).
				ExcludeByField(config.ExcludeFieldPatterns).
				FilterByNotHasDirective(config.Directive).
				FilterByPositionNotNil()

			for _, field := range fields {
				pass.Reportf(field.Position, config.ReportFormat, field.Name)
			}

			return nil, nil
		},
	}
}

func ArgumentAnalyzer(config *ArgumentConfig) *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: config.Directive,
		Doc:  config.Description,
		Run: func(pass *gqlanalysis.Pass) (any, error) {
			definitions := NewDefinitionsByMap(pass.Schema.Types).NotBuildIn().FilterByKinds(config.Kinds)
			fields := definitions.Fields().FilterByNotNil()
			arguments := fields.Arguments().
				FilterByArgumentType(config.ArgumentTypePatterns).
				ExcludeByArgumentName(config.ExcludeArgumentPatterns).
				FilterByNotHasDirective(config.Directive).
				FilterByPositionNotNil()

			for _, argument := range arguments {
				pass.Reportf(argument.Position, config.ReportFormat, argument.Name)
			}
			return nil, nil
		},
	}
}
