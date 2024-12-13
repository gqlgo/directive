package directive

import (
	"github.com/vektah/gqlparser/v2/ast"
	"slices"
)

type ArgumentDefinitions ast.ArgumentDefinitionList

func (as ArgumentDefinitions) FilterByArgumentType(argumentTypePatterns []string) ArgumentDefinitions {
	var arguments ArgumentDefinitions
	for _, argument := range as {
		if slices.ContainsFunc(argumentTypePatterns, func(pattern string) bool { return isType(pattern, argument.Type) }) {
			arguments = append(arguments, argument)
		}
	}

	return arguments
}

func (as ArgumentDefinitions) ExcludeByArgumentName(excludeArgumentPatterns []string) ArgumentDefinitions {
	var arguments ArgumentDefinitions
	for _, argument := range as {
		if !slices.ContainsFunc(excludeArgumentPatterns, func(pattern string) bool { return isMatch(pattern, argument.Name) }) {
			arguments = append(arguments, argument)
		}
	}
	return arguments
}

func (as ArgumentDefinitions) FilterByNotHasDirective(directive string) ArgumentDefinitions {
	var arguments ArgumentDefinitions
	for _, argument := range as {
		if !findDirectiveOnArgument(argument, directive) {
			arguments = append(arguments, argument)
		}
	}
	return arguments
}

func (as ArgumentDefinitions) FilterByPositionNotNil() ArgumentDefinitions {
	var arguments ArgumentDefinitions
	for _, argument := range as {
		if argument.Position != nil {
			arguments = append(arguments, argument)
		}
	}
	return arguments
}

func findDirectiveOnArgument(argument *ast.ArgumentDefinition, directiveName string) bool {
	return argument.Directives.ForName(directiveName) != nil
}
