package goqumysqllint

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("goqumysqllint", newPlugin)
}

func newPlugin(_ any) (register.LinterPlugin, error) {
	return &Plugin{}, nil
}

type Plugin struct{}

func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	analyzer, err := NewAnalyzer(nil)
	if err != nil {
		return nil, err
	}
	return []*analysis.Analyzer{analyzer}, nil
}

func (p *Plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
