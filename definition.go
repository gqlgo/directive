package directive

import (
	"github.com/vektah/gqlparser/v2/ast"
	"maps"
	"slices"
)

type Definitions ast.DefinitionList

func NewDefinitionsByMap(typeByName map[string]*ast.Definition) Definitions {
	return slices.Collect(maps.Values(typeByName))
}

func (ds Definitions) Fields() FieldDefinitions {
	var fields FieldDefinitions
	for _, definition := range ds {
		fields = append(fields, definition.Fields...)
	}
	return fields
}

func (ds Definitions) NotBuildIn() Definitions {
	var definitions Definitions
	for _, definition := range ds {
		if !definition.BuiltIn {
			definitions = append(definitions, definition)
		}
	}
	return definitions
}

func (ds Definitions) FilterByKinds(kinds []ast.DefinitionKind) Definitions {
	var definitions Definitions
	for _, definition := range ds {
		if slices.Contains(kinds, definition.Kind) {
			definitions = append(definitions, definition)
		}
	}
	return definitions
}

func (ds Definitions) FilterByDefinitionName(definitionPattern []string) Definitions {
	var definitions Definitions
	for _, definition := range ds {
		if slices.ContainsFunc(definitionPattern, func(pattern string) bool { return isMatch(pattern, definition.Name) }) {
			definitions = append(definitions, definition)
		}
	}
	return definitions
}

func (ds Definitions) ExcludeByDefinitionName(excludeDefinitionPatterns []string) Definitions {
	var definitions Definitions
	for _, definition := range ds {
		if !slices.ContainsFunc(excludeDefinitionPatterns, func(pattern string) bool { return isMatch(pattern, definition.Name) }) {
			definitions = append(definitions, definition)
		}
	}
	return definitions
}

func (ds Definitions) FilterByNotHasDirective(directive string) Definitions {
	var definitions Definitions
	for _, definition := range ds {
		if !findDirectiveOnDefinition(definition, directive) {
			definitions = append(definitions, definition)
		}
	}
	return definitions
}

func (ds Definitions) FilterByPositionNotNil() Definitions {
	var definitions Definitions
	for _, definition := range ds {
		if definition.Position != nil {
			definitions = append(definitions, definition)
		}
	}
	return definitions
}

func findDirectiveOnDefinition(definition *ast.Definition, directiveName string) bool {
	return definition.Directives.ForName(directiveName) != nil
}
