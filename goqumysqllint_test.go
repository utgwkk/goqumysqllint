package goqumysqllint

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	t.Run("boolean_comparison", func(t *testing.T) {
		analyzer, err := NewAnalyzer(nil)
		if err != nil {
			t.Fatal(err)
		}
		testdata := testutil.WithModules(t, analysistest.TestData(), nil)
		analysistest.Run(t, testdata, analyzer, "boolean_comparison")
	})
}
