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
	for _, c := range analyzerConfig.FieldConfig {
		analyzer := FieldAnalyzer(c)
		analyzers = append(analyzers, analyzer)
	}
	for _, c := range analyzerConfig.ArgumentConfig {
		analyzer := ArgumentAnalyzer(c)
		analyzers = append(analyzers, analyzer)
	}
	for _, c := range analyzerConfig.TypeConfig {
		analyzer := TypeAnalyzer(c)
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

func FieldAnalyzer(config *FieldConfig) *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: config.Directive,
		Doc:  config.Description,
		Run:  FieldLinter(config),
	}
}

func FieldLinter(config *FieldConfig) func(pass *gqlanalysis.Pass) (any, error) {
	return func(pass *gqlanalysis.Pass) (any, error) {
		types := targetTypeKind(pass.Schema.Types, config.Kinds)
		types2 := targetTypes(types, config.FieldParentTypePatterns)
		for _, t := range types2 {
			fields := targetFieldType(t.Fields, config.FieldTypePatterns)
			excludedFields := excludeTargetFieldTypeByTypeName(fields, config.IgnoreFieldPatterns)
			for _, field := range excludedFields {
				if !findDirectiveOnField(field, config.Directive) {
					if field.Position != nil {
						pass.Reportf(field.Position, config.ReportFormat, t.Name, field.Name)
					}
				}
			}
		}
		return nil, nil
	}
}

func ArgumentAnalyzer(config *ArgumentConfig) *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: config.Directive,
		Doc:  config.Description,
		Run:  argumentAnalyzer(config),
	}
}

func argumentAnalyzer(config *ArgumentConfig) func(pass *gqlanalysis.Pass) (any, error) {
	return func(pass *gqlanalysis.Pass) (any, error) {
		types := targetTypeKind(pass.Schema.Types, config.Kinds)
		for _, t := range types {
			for _, field := range t.Fields {
				args := targetFieldArgumentType(field, config.ArgumentTypePatterns)
				excludedArgs := excludeTargetArgumentsByField(args, config.IgnoreArgumentPatterns)
				for _, arg := range excludedArgs {
					if !findDirectiveOnArg(arg, config.Directive) {
						if arg.Position != nil {
							pass.Reportf(arg.Position, config.ReportFormat, arg.Name, field.Name)
						}
					}
				}
			}
		}
		return nil, nil
	}
}

func TypeAnalyzer(config *TypeConfig) *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: config.Directive,
		Doc:  config.Description,
		Run:  typeAnalyzer(config),
	}
}

func typeAnalyzer(config *TypeConfig) func(pass *gqlanalysis.Pass) (any, error) {
	return func(pass *gqlanalysis.Pass) (any, error) {
		types := targetTypeKind(pass.Schema.Types, config.Kinds)
		excludedTypes := excludeTargetTypesByTypeName(types, config.IgnoreTypePatterns)
		for _, t := range excludedTypes {
			if !findDirectiveOnDefinition(t, config.Directive) {
				if t.Position != nil {
					pass.Reportf(t.Position, config.ReportFormat, t.Name)
				}
			}
		}
		return nil, nil
	}
}
