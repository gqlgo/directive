package directive

import (
	"github.com/gqlgo/gqlanalysis"
	"github.com/vektah/gqlparser/v2/ast"
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
	for _, c := range analyzerConfig.InputObjectFieldConfig {
		analyzer := InputObjectFieldAnalyzer(c)
		analyzers = append(analyzers, analyzer)
	}
	for _, c := range analyzerConfig.ObjectFieldArgumentConfig {
		analyzer := ObjectFieldArgumentAnalyzer(c)
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

func InputObjectFieldAnalyzer(config *InputObjectFieldConfig) *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: config.Directive,
		Doc:  config.Description,
		Run:  inputObjectFieldLinter(config),
	}
}

func inputObjectFieldLinter(config *InputObjectFieldConfig) func(pass *gqlanalysis.Pass) (any, error) {
	return func(pass *gqlanalysis.Pass) (any, error) {
		types := targetTypes(pass.Schema.Types, []ast.DefinitionKind{ast.InputObject})
		for _, t := range types {
			fields := targetFieldType(t.Fields, config.FieldTypePatterns)
			excludedFields := excludeTargetFieldTypeByTypeName(fields, config.IgnoreFieldNamePatterns)
			for _, field := range excludedFields {
				if !findDirectiveOnField(field, config.Directive) {
					pass.Reportf(field.Position, config.ReportFormat, t.Name, field.Name)
				}
			}
		}
		return nil, nil
	}
}

func ObjectFieldArgumentAnalyzer(config *ObjectFieldArgumentConfig) *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: config.Directive,
		Doc:  config.Description,
		Run:  objectFieldArgumentAnalyzer(config),
	}
}

func objectFieldArgumentAnalyzer(config *ObjectFieldArgumentConfig) func(pass *gqlanalysis.Pass) (any, error) {
	return func(pass *gqlanalysis.Pass) (any, error) {
		types := targetTypes(pass.Schema.Types, []ast.DefinitionKind{ast.Object})
		for _, t := range types {
			for _, field := range t.Fields {
				args := targetFieldArgumentType(field, config.ArgumentTypePatterns)
				excludedArgs := excludeTargetArgumentsByFieldName(args, config.IgnoreArgumentNamePatterns)
				for _, arg := range excludedArgs {
					if !findDirectiveOnArg(arg, config.Directive) {
						pass.Reportf(arg.Position, config.ReportFormat, arg.Name, field.Name)
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
		types := targetTypes(pass.Schema.Types, config.Kinds)
		excludedTypes := excludeTargetTypesByTypeName(types, config.IgnoreTypePatterns)
		for _, t := range excludedTypes {
			if !findDirectiveOnDefinition(t, config.Directive) {
				pass.Reportf(t.Position, config.ReportFormat, t.Name)
			}
		}
		return nil, nil
	}
}
