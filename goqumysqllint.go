package goqumysqllint

import (
	"go/ast"
	"strings"

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

type analyzer struct{}

func newAnalyzer() (*analyzer, error) {
	return &analyzer{}, nil
}

func (a *analyzer) run(pass *analysis.Pass) (interface{}, error) {
	it := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector).PreorderSeq(
		new(ast.CallExpr),
		new(ast.CompositeLit),
	)

	for node := range it {
		switch node := node.(type) {
		case *ast.CallExpr:
			a.checkBooleanExpressionComparison(pass, node)
		case *ast.CompositeLit:
			a.checkGoquExBooleanComparison(pass, node)
		}
	}
	return nil, nil
}

func (a *analyzer) checkBooleanExpressionComparison(pass *analysis.Pass, node *ast.CallExpr) {
	ty := pass.TypesInfo.TypeOf(node)
	if !strings.HasSuffix(ty.String(), "exp.BooleanExpression") {
		return
	}
	if len(node.Args) != 1 {
		return
	}
	arg := node.Args[0]
	argTy := pass.TypesInfo.TypeOf(arg)
	if argTy.String() != "bool" {
		return
	}

	pass.Reportf(node.Pos(), "compare boolean value with int")
}

func (a *analyzer) checkGoquExBooleanComparison(pass *analysis.Pass, node *ast.CompositeLit) {
	selector, ok := node.Type.(*ast.SelectorExpr)
	if !ok {
		return
	}
	xIdent, ok := selector.X.(*ast.Ident)
	if !ok {
		return
	}
	if !(xIdent.Name == "goqu" || xIdent.Name == "exp") {
		return
	}
	if !(selector.Sel.Name == "Ex" || selector.Sel.Name == "ExOr") {
		return
	}
	for _, elt := range node.Elts {
		kvExpr, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		v := kvExpr.Value
		vTy := pass.TypesInfo.TypeOf(v)
		if vTy.String() != "bool" {
			continue
		}

		pass.Reportf(node.Pos(), "compare boolean value with int")
	}
}
