package main

import (
	"github.com/gqlgo/gqlanalysis/multichecker"
	"github.com/gqlgo/iddirective"
)

func main() {
	multichecker.Main(
		iddirective.Analyzer(),
	)
}
