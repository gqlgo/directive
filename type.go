package directive

import (
	"fmt"
	"github.com/vektah/gqlparser/v2/ast"
	"slices"
)

type Types []*ast.Type

func (ts Types) FilterByType(typePatterns []string) Types {
	var types Types
	for _, t := range ts {
		if slices.ContainsFunc(typePatterns, func(pattern string) bool { return isType(pattern, t) }) {
			types = append(types, t)
		}
	}

	return types
}

func targetFieldType(fields ast.FieldList, fieldTypePatterns []string) ast.FieldList {
	var targets ast.FieldList
	for _, field := range fields {
		if slices.ContainsFunc(fieldTypePatterns, func(pattern string) bool { return isType(pattern, field.Type) }) {
			targets = append(targets, field)
		}
	}
	return targets
}

func isType(pattern string, t *ast.Type) bool {
	if t == nil {
		return false
	}
	return isMatch(pattern, typeName(t))
}

func typeName(t *ast.Type) string {
	if isList(t) {
		return fmt.Sprintf("[%s]", t.Elem.NamedType)
	}
	return t.NamedType
}

func isList(t *ast.Type) bool {
	return t.Elem != nil
}
