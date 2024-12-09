package directive

import (
	"github.com/vektah/gqlparser/v2/ast"
)

func targetTypes(types map[string]*ast.Definition, kind ast.DefinitionKind) []*ast.Definition {
	var targets []*ast.Definition
	for _, t := range types {
		if !t.BuiltIn && t.Kind == kind {
			targets = append(targets, t)
		}
	}
	return targets
}

func targetFieldType(fields ast.FieldList, fieldType string) ast.FieldList {
	var targets ast.FieldList
	for _, field := range fields {
		if field != nil && field.Type != nil {
			if isType(field.Type, fieldType) {
				targets = append(targets, field)
			}
		}
	}
	return targets
}

func targetFieldArgumentType(field *ast.FieldDefinition, fieldArgumentType string) ast.ArgumentDefinitionList {
	var targets ast.ArgumentDefinitionList
	if field != nil && field.Type != nil {
		for _, arg := range field.Arguments {
			if isType(arg.Type, fieldArgumentType) {
				targets = append(targets, arg)
			}
		}
	}
	return targets
}

func findDirectiveOnField(field *ast.FieldDefinition, directiveName string) bool {
	return field.Directives.ForName(directiveName) != nil
}

func findDirectiveOnArg(arg *ast.ArgumentDefinition, directiveName string) bool {
	return arg.Directives.ForName(directiveName) != nil
}

func isType(t *ast.Type, typeName string) bool {
	if t == nil {
		return false
	}
	if t.NamedType == typeName {
		return true
	}
	return isType(t.Elem, typeName)
}
