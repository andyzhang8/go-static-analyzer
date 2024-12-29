package rules

import (
	"go-static-analyzer/report"
	"go-static-analyzer/utils"
	"go/ast"
	"go/token"
)

// Max allowed lines in a function
const MaxFunctionLength = 50

// Ensures functions are not excessively long
func CheckFunctionLength(node *ast.File, fset *token.FileSet, filePath string) {
	ast.Inspect(node, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok {
			lines := utils.GetLine(fn.Body.End(), fset) - utils.GetLine(fn.Pos(), fset)
			if lines > MaxFunctionLength {
				report.AddIssue(filePath, "Excessive Function Length", utils.GetLine(fn.Pos(), fset), 0, "Function '"+fn.Name.Name+"' exceeds "+utils.IntToString(MaxFunctionLength)+" lines")
			}
		}
		return true
	})
}
