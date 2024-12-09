package directive

import (
	"github.com/gqlgo/gqlanalysis"
)

func InputObjectFieldAnalyzer(config *InputObjectFieldConfig) *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: "directive",
		Doc:  "detect missing directives.",
		Run:  inputObjectFieldLinter(config),
	}
}

// inputObjectFieldLinter detects missing directives for input object field.
func inputObjectFieldLinter(config *InputObjectFieldConfig) func(pass *gqlanalysis.Pass) (interface{}, error) {
	return func(pass *gqlanalysis.Pass) (interface{}, error) {
		types := targetTypes(pass.Schema.Types, config.Kind)
		for _, t := range types {
			fields := targetFieldType(t.Fields, config.FieldType)
			for _, field := range fields {
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
		Name: "directive",
		Doc:  "detect missing directives.",
		Run:  objectFieldArgumentAnalyzer(config),
	}
}

// inputObjectFieldLinter detects missing directives for input object field.
func objectFieldArgumentAnalyzer(config *ObjectFieldArgumentConfig) func(pass *gqlanalysis.Pass) (interface{}, error) {
	return func(pass *gqlanalysis.Pass) (interface{}, error) {
		types := targetTypes(pass.Schema.Types, config.Kind)
		for _, t := range types {
			for _, field := range t.Fields {
				args := targetFieldArgumentType(field, config.FieldArgumentType)
				for _, arg := range args {
					if !findDirectiveOnArg(arg, config.Directive) {
						pass.Reportf(field.Position, config.ReportFormat, arg.Name, field.Name)
					}
				}
			}
		}
		return nil, nil
	}
}
