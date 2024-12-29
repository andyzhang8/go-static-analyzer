package rules

import (
	"go-static-analyzer/report"
	"go-static-analyzer/utils"
	"go/ast"
	"go/token"
)

// Detects nested loops
func CheckNestedLoops(node *ast.File, fset *token.FileSet, filePath string) {
	ast.Inspect(node, func(n ast.Node) bool {
		if forStmt, ok := n.(*ast.ForStmt); ok {
			ast.Inspect(forStmt.Body, func(innerNode ast.Node) bool {
				if _, nestedOk := innerNode.(*ast.ForStmt); nestedOk {
					report.AddIssue(filePath, "Nested Loops", utils.GetLine(forStmt.Pos(), fset), 0, "Avoid nested loops for better performance")
				}
				return true
			})
		}
		return true
	})
}
