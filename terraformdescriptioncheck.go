package terraformdescriptioncheck

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "terraformdescriptioncheck",
	Doc:  "checks that *schema.Schema has non-empty Description",
	Run:  run,
}

func NewAnalyzer() *analysis.Analyzer {
	return Analyzer
}

func run(pass *analysis.Pass) (interface{}, error) {
	ast.Inspect(pass.Files[0], func(n ast.Node) bool {
		composite, ok := n.(*ast.CompositeLit)
		if !ok {
			return true
		}

		// Проверка типа: *schema.Schema
		selExpr, ok := composite.Type.(*ast.SelectorExpr)
		if !ok || selExpr.Sel.Name != "Schema" {
			return true
		}

		ident, ok := selExpr.X.(*ast.Ident)
		if !ok || ident.Name != "schema" {
			return true
		}

		for _, elt := range composite.Elts {
			kv, ok := elt.(*ast.KeyValueExpr)
			if !ok {
				continue
			}

			keyIdent, ok := kv.Key.(*ast.Ident)
			if !ok || keyIdent.Name != "Description" {
				continue
			}

			val, ok := kv.Value.(*ast.BasicLit)
			if !ok || val.Kind != token.STRING {
				continue
			}

			if val.Value == `""` {
				pass.Reportf(val.Pos(), "Description in schema.Schema is empty")
			}
		}

		return true
	})

	return nil, nil
}
