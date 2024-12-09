# directive

[![pkg.go.dev][gopkg-badge]][gopkg]

`directive` finds id fields with no @id directive and arguments in your GraphQL schema files.

```graphql
input NoIdDirectiveMutationInput {
    name: String!
    adminID: ID! # want "adminID has no id directive"
}
```

## How to use

A runnable linter can be created with multichecker package.
You can create own linter with your favorite Analyzers.

```go
package main

import (
	"flag"
	"github.com/gqlgo/directive"
	"github.com/gqlgo/gqlanalysis/multichecker"
)

func main() {
	multichecker.Main(
		directive.Analyzer(),
	)
}
```

`directive` provides a typical main function and you can install with `go install` command.

```sh
$ go install github.com/gqlgo/directive/cmd/directive@latest
```

The `directive` command has a flag, `schema` which will be parsed and analyzed by directive's Analyzer.

```sh
$ directive -schema="server/graphql/schema/**/*.graphql"
```

The default value of `schema` is "schema/*/**.graphql".

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gqlgo/directive
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gqlgo/directive?status.svg
