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
    input_object_field:
      - description: "constraint directive exists on the input field"
        directive: constraint
        field_type: ['^\[?Int\]?$', '^\[?Float\]?$', '^\[?String\]?$', '^\[?Decimal\]?$', '^\[?URL\]?$']
        ignore_field_name: ['^first$', '^last$', '^after$', '^before$']
        report_format: "%s.%s has no constraint directive"
    object_field_argument:
      - description: "constraint directive exists on the object field argument"
        directive: constraint
        argument_type: ['^\[?Int\]?$', '^\[?Float\]?$', '^\[?String\]?$', '^\[?Decimal\]?$', '^\[?URL\]?$']
        ignore_argument_name: ['^first$', '^last$', '^after$', '^before$']
        report_format: "argument %s of %s has no constraint directive"
```

- sample2: permission directive exists on the type

```yaml
---
analyzer:
  - analyzer_name: "permission directive"
    description: "permission directive exists on the type"
    type:
      - description: "permission directive exists on the type"
        directive: permission
        kind: ['OBJECT', 'INTERFACE', 'FIELD_DEFINITION']
        object_patterns: ['.*']
        ignore_type: [ '^Query$', '^Mutation$', '^Subscription$', 'PageInfo$', 'Connection$', 'Payload$']
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
