package iddirective_test

import (
	"testing"

	"github.com/gqlgo/gqlanalysis/analysistest"
	"github.com/gqlgo/iddirective"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData(t)
	analysistest.Run(t, testdata, iddirective.Analyzer(), "a")
}
