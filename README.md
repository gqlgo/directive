# directive

[![pkg.go.dev][gopkg-badge]][gopkg]

`directive` finds id fields with no @id directive and arguments in your GraphQL schema files.

```graphql
input UserCreateInput {
    name: String! # want "UserCreateInput.name has no constraint directive"
}
```

## Usage

`directive` provides a typical main function and you can install with `go install` command.

```sh
$ go install github.com/gqlgo/directive/cmd/directive@latest
```

The `directive` command has a flag, `config` which will be parsed and analyzed by the Analyzer.


- sample1: constraint directive exists on the field

```yaml
---
analyzer:
  - analyzer_name: "constraint directive"
    description: "constraint directive exists on the field"
    field:
      - description: "constraint directive exists on the input field"
        directive: constraint
        kind: ['INPUT_OBJECT']
        field_parent_type: ['.+']
        field_type: ['^\[?Int\]?$', '^\[?Float\]?$', '^\[?String\]?$', '^\[?Decimal\]?$', '^\[?URL\]?$']
        exclude_field: ['^first$', '^last$', '^after$', '^before$']
        report_format: "%s has no constraint directive"
    argument:
      - description: "constraint directive exists on the object field argument"
        directive: constraint
        kind: ['OBJECT']
        argument_type: ['^\[?Int\]?$', '^\[?Float\]?$', '^\[?String\]?$', '^\[?Decimal\]?$', '^\[?URL\]?$']
        exclude_argument: ['^first$', '^last$', '^after$', '^before$']
        report_format: "argument %s has no constraint directive"
```

- sample2: permission directive exists on the definition

```yaml
---
analyzer:
  - analyzer_name: "permission directive"
    description: "permission directive exists on the definition"
    definition:
      - description: "permission directive exists on the definition"
        directive: permission
        kind: ['OBJECT', 'INTERFACE']
        definition: ['.+']
        exclude_definition: [ '^Query$', '^Mutation$', '^Subscription$', '^PageInfo$']
        report_format: "%s has no permission directive"
    field:
      - description: "permission directive exists on the mutation"
        directive: permission
        kind: ['OBJECT']
        field_parent_type:  ['^Mutation$']
        field_type: ['.+']
        exclude_field:
        report_format: "%s has no permission directive"

```

The `directive` command has a flag, `schema` which will be parsed and analyzed by directive's Analyzer.

```sh
$ directive -schema="server/graphql/schema/**/*.graphql"
```

The default value of `schema` is "schema/*/**.graphql".

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gqlgo/directive
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gqlgo/directive?status.svg
