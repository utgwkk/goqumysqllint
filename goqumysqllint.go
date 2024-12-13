package goqumysqllint

import (
	"errors"
	"fmt"
	"go/ast"
	"go/types"

	_ "github.com/doug-martin/goqu/v9" // for lookup types
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/packages"
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
	it := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector).PreorderSeq()
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.LoadTypes | packages.NeedTypesInfo,
	}, "github.com/doug-martin/goqu/v9/exp")
	if err != nil {
		return nil, err
	}
	goquPkg := pkgs[0]
	boolExpTy := goquPkg.Types.Scope().Lookup("BooleanExpression").Type().Underlying().(*types.Interface)
	if boolExpTy == nil {
		return nil, errors.New("BooleanExpression not found")
	}
	fmt.Println(boolExpTy)
	for node := range it {
		node, ok := node.(ast.Expr)
		if !ok {
			continue
		}
		ty := pass.TypesInfo.TypeOf(node)
		fmt.Println(ty)
		fmt.Println(types.Implements(ty, boolExpTy))
	}
	return nil, nil
}
