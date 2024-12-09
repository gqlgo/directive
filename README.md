# iddirective

[![pkg.go.dev][gopkg-badge]][gopkg]

`iddirective` finds id fields with no @id directive and arguments in your GraphQL schema files.

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
	"github.com/gqlgo/iddirective"
	"github.com/gqlgo/gqlanalysis/multichecker"
)

func main() {
	multichecker.Main(
		iddirective.Analyzer(),
	)
}
```

`iddirective` provides a typical main function and you can install with `go install` command.

```sh
$ go install github.com/gqlgo/iddirective/cmd/iddirective@latest
```

The `iddirective` command has a flag, `schema` which will be parsed and analyzed by iddirective's Analyzer.

```sh
$ iddirective -schema="server/graphql/schema/**/*.graphql"
```

The default value of `schema` is "schema/*/**.graphql".

## Author

[![Appify Technologies, Inc.](appify-logo.png)](http://github.com/appify-technologies)

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gqlgo/iddirective
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gqlgo/iddirective?status.svg
