package iddirective

import (
	"github.com/gqlgo/gqlanalysis"
	"github.com/vektah/gqlparser/v2/ast"
)

func Analyzer() *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: "iddirective",
		Doc:  "iddirective finds id fields with no id directive.",
		Run:  run,
	}
}

func isIDType(t *ast.Type) bool {
	if t == nil {
		return false
	}
	if t.NamedType == "ID" {
		return true
	}
	return isIDType(t.Elem)
}

func run(pass *gqlanalysis.Pass) (interface{}, error) {
	for _, t := range pass.Schema.Types {
		if t.BuiltIn {
			continue
		}
		if t.Kind == ast.InputObject {
			for _, field := range t.Fields {
				if field != nil && field.Type != nil {
					if isIDType(field.Type) {
						if field.Directives.ForName("id") == nil {
							pass.Reportf(field.Position, "%s has no id directive", field.Name)
						}
					}
				}
			}
		}
		if t.Kind == ast.Object {
			for _, field := range t.Fields {
				for _, arg := range field.Arguments {
					if isIDType(arg.Type) {
						if arg.Directives.ForName("id") == nil {
							pass.Reportf(field.Position, "argument %s of %s has no id directive", arg.Name, field.Name)
						}
					}
				}
			}
		}
	}

	return nil, nil
}
