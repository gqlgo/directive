package directive_test

import (
	"testing"

	"github.com/gqlgo/directive"
	"github.com/gqlgo/gqlanalysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData(t)
	analysistest.Run(t, testdata, directive.Analyzer(), "a")
}
