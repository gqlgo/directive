package directive

import (
	"fmt"
	"github.com/vektah/gqlparser/v2/ast"
	"regexp"
	"slices"
)

func targetTypeKind(types map[string]*ast.Definition, kinds []ast.DefinitionKind) []*ast.Definition {
	var targets []*ast.Definition
	for _, t := range types {
		if !t.BuiltIn && slices.Contains(kinds, t.Kind) {
			targets = append(targets, t)
		}
	}
	return targets
}

func targetTypes(types ast.DefinitionList, fieldParentTypePatterns []string) ast.DefinitionList {
	var targets ast.DefinitionList
	for _, t := range types {
		if slices.ContainsFunc(fieldParentTypePatterns, func(pattern string) bool { return isNameMatch(pattern, t.Name) }) {
			targets = append(targets, t)
		}
	}
	return targets
}

func targetFieldType(fields ast.FieldList, fieldTypePatterns []string) ast.FieldList {
	var targets ast.FieldList
	for _, field := range fields {
		if field != nil && field.Type != nil {
			if slices.ContainsFunc(fieldTypePatterns, func(pattern string) bool { return isType(pattern, field.Type) }) {
				targets = append(targets, field)
			}
		}
	}
	return targets
}

func excludeTargetFieldTypeByTypeName(fields ast.FieldList, ignoreFiledNamePatterns []string) ast.FieldList {
	var targets ast.FieldList
	for _, field := range fields {
		if field != nil && field.Type != nil {
			if !slices.ContainsFunc(ignoreFiledNamePatterns, func(pattern string) bool { return isNameMatch(pattern, field.Name) }) {
				targets = append(targets, field)
			}
		}
	}
	return targets
}

func targetFieldArgumentType(field *ast.FieldDefinition, fieldArgumentTypePatterns []string) ast.ArgumentDefinitionList {
	var targets ast.ArgumentDefinitionList
	if field != nil && field.Type != nil {
		for _, arg := range field.Arguments {
			if slices.ContainsFunc(fieldArgumentTypePatterns, func(pattern string) bool { return isType(pattern, arg.Type) }) {
				targets = append(targets, arg)
			}
		}
	}
	return targets
}

func excludeTargetArgumentsByField(args ast.ArgumentDefinitionList, ignoreArgumentPatterns []string) ast.ArgumentDefinitionList {
	var targets ast.ArgumentDefinitionList
	for _, arg := range args {
		if !slices.ContainsFunc(ignoreArgumentPatterns, func(pattern string) bool { return isNameMatch(pattern, arg.Name) }) {
			targets = append(targets, arg)
		}
	}
	return targets
}

func excludeTargetTypesByTypeName(types ast.DefinitionList, ignoreTypePatterns []string) ast.DefinitionList {
	var targets ast.DefinitionList
	for _, t := range types {
		if !slices.ContainsFunc(ignoreTypePatterns, func(pattern string) bool { return isNameMatch(pattern, t.Name) }) {
			targets = append(targets, t)
		}
	}
	return targets
}

func findDirectiveOnDefinition(t *ast.Definition, directiveName string) bool {
	return t.Directives.ForName(directiveName) != nil
}

func findDirectiveOnField(field *ast.FieldDefinition, directiveName string) bool {
	return field.Directives.ForName(directiveName) != nil
}

func findDirectiveOnArg(arg *ast.ArgumentDefinition, directiveName string) bool {
	return arg.Directives.ForName(directiveName) != nil
}

func isType(pattern string, t *ast.Type) bool {
	if t == nil {
		return false
	}
	return isNameMatch(pattern, typeName(t))
}

func isNameMatch(pattern, name string) bool {
	return regexp.MustCompile(pattern).MatchString(name)
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
