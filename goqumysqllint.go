package goqumysqllint

import (
	"go/ast"

	_ "github.com/doug-martin/goqu/v9" // for lookup types
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
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
		Requires: []*analysis.Analyzer{
			inspect.Analyzer,
		},
	}, nil
}

const (
	booleanExpressionFullTypeName = "github.com/doug-martin/goqu/v9/exp.BooleanExpression"
)

type analyzer struct{}

func newAnalyzer() (*analyzer, error) {
	return &analyzer{}, nil
}

func (a *analyzer) run(pass *analysis.Pass) (interface{}, error) {
	it := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector).PreorderSeq(
		new(ast.CallExpr),
	)

	for node := range it {
		node, ok := node.(*ast.CallExpr)
		if !ok {
			continue
		}
		ty := pass.TypesInfo.TypeOf(node)
		if ty.String() != booleanExpressionFullTypeName {
			continue
		}
		if len(node.Args) != 1 {
			continue
		}
		arg := node.Args[0]
		argTy := pass.TypesInfo.TypeOf(arg)
		if argTy.String() != "bool" {
			continue
		}

		pass.Report(analysis.Diagnostic{
			Pos:     node.Pos(),
			Message: "compare boolean value with int",
		})
	}
	return nil, nil
}
