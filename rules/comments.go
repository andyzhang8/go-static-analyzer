package rules

import (
	"go-static-analyzer/report"
	"go-static-analyzer/utils"
	"go/ast"
	"go/token"
)

// Ensures top-level functions have comments
func CheckComments(node *ast.File, fset *token.FileSet, filePath string) {
	ast.Inspect(node, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok && fn.Doc == nil {
			report.AddIssue(filePath, "Missing Comments", utils.GetLine(fn.Pos(), fset), 0, "Function '"+fn.Name.Name+"' is missing a comment")
		}
		return true
	})
}
