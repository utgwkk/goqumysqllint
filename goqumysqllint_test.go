package goqumysqllint

import (
	"strings"
	"testing"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

const dummyGoModContent = `module dummy

go 1.23.4
`

func TestAnalyzer(t *testing.T) {
	t.Run("boolean_comparison", func(t *testing.T) {
		analyzer, err := NewAnalyzer(nil)
		if err != nil {
			t.Fatal(err)
		}
		gomod := strings.NewReader(dummyGoModContent)
		testdata := testutil.WithModules(t, analysistest.TestData(), gomod)
		analysistest.Run(t, testdata, analyzer, "boolean_comparison")
	})
}
