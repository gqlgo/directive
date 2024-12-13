package directive

import (
	"github.com/vektah/gqlparser/v2/ast"
	"slices"
)

type FieldDefinitions ast.FieldList

func (fs FieldDefinitions) Arguments() ArgumentDefinitions {
	var arguments ArgumentDefinitions
	for _, field := range fs {
		arguments = append(arguments, field.Arguments...)
	}
	return arguments
}

func (fs FieldDefinitions) FilterByNotNil() FieldDefinitions {
	var fields FieldDefinitions
	for _, field := range fs {
		if field != nil && field.Type != nil {
			fields = append(fields, field)
		}
	}

	return fields
}
func (fs FieldDefinitions) FilterByFieldType(fieldTypePatterns []string) FieldDefinitions {
	var fields FieldDefinitions
	for _, field := range fs {
		if slices.ContainsFunc(fieldTypePatterns, func(pattern string) bool { return isType(pattern, field.Type) }) {
			fields = append(fields, field)
		}
	}

	return fields
}

func (fs FieldDefinitions) ExcludeByField(excludeFiledPatterns []string) FieldDefinitions {
	var fields FieldDefinitions
	for _, field := range fs {
		if !slices.ContainsFunc(excludeFiledPatterns, func(pattern string) bool { return isMatch(pattern, field.Name) }) {
			fields = append(fields, field)
		}
	}

	return fields
}

func (fs FieldDefinitions) FilterByNotHasDirective(directive string) FieldDefinitions {
	var fields FieldDefinitions
	for _, field := range fs {
		if !findDirectiveOnField(field, directive) {
			if field.Position != nil {
				fields = append(fields, field)
			}
		}
	}
	return fields
}

func (fs FieldDefinitions) FilterByPositionNotNil() FieldDefinitions {
	var fields FieldDefinitions
	for _, field := range fs {
		if field.Position != nil {
			fields = append(fields, field)
		}
	}
	return fields
}

func findDirectiveOnField(field *ast.FieldDefinition, directiveName string) bool {
	return field.Directives.ForName(directiveName) != nil
}
