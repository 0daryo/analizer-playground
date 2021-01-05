package mylinter

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "mylinter",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

const doc = "unused find unused identifyers"

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.BinaryExpr)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		e := n.(*ast.BinaryExpr)

		// Only want == or != comparisons.
		if e.Op != token.EQL && e.Op != token.NEQ {
			return
		}

		// Only want comparisons with a nil identifier on one side.
		var e2 ast.Expr
		switch {
		case pass.TypesInfo.Types[e.X].IsNil():
			e2 = e.Y
		case pass.TypesInfo.Types[e.Y].IsNil():
			e2 = e.X
		default:
			return
		}

		// Only want identifiers or selector expressions.
		var obj types.Object
		switch v := e2.(type) {
		case *ast.Ident:
			obj = pass.TypesInfo.Uses[v]
		case *ast.SelectorExpr:
			obj = pass.TypesInfo.Uses[v.Sel]
		default:
			return
		}

		// Only want functions.
		if _, ok := obj.(*types.Func); !ok {
			return
		}

		pass.ReportRangef(e, "comparison of function %v %v nil is always %v", obj.Name(), e.Op, e.Op == token.NEQ)
	})
	return nil, nil
}
