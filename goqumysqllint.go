package goqumysqllint

import (
	"golang.org/x/tools/go/analysis"
)

func NewAnalyzer(_ any) (*analysis.Analyzer, error) {
	anylizer, err := newAnalyzer()
	if err != nil {
		return nil, err
	}
	return &analysis.Analyzer{
		Name: "goqumysqllint",
		Doc:  "golang mysql linter",
		Run:  anylizer.run,
	}, nil
}

type analyzer struct{}

func newAnalyzer() (*analyzer, error) {
	return &analyzer{}, nil
}

func (a *analyzer) run(pass *analysis.Pass) (interface{}, error) {
	return nil, nil
}
