package main

import (
	"github.com/gqlgo/directive"
	"github.com/gqlgo/gqlanalysis/multichecker"
)

func main() {
	multichecker.Main(
		directive.Analyzer(),
	)
}
