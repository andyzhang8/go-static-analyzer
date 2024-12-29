package rules

import (
	"go-static-analyzer/report"
	"go-static-analyzer/utils"
	"go/ast"
	"go/token"
	"strings"
)

// Detects hardcoded sensitive strings
func CheckHardcodedCredentials(node *ast.File, fset *token.FileSet, filePath string) {
	ast.Inspect(node, func(n ast.Node) bool {
		if assign, ok := n.(*ast.AssignStmt); ok {
			for _, expr := range assign.Rhs {
				if lit, ok := expr.(*ast.BasicLit); ok && lit.Kind.String() == "STRING" {
					cleaned := utils.TrimQuotes(lit.Value)
					if strings.Contains(cleaned, "password") || strings.Contains(cleaned, "api_key") {
						report.AddIssue(filePath, "Hardcoded Credentials", utils.GetLine(assign.Pos(), fset), 0, "Potential hardcoded credential: "+cleaned)
					}
				}
			}
		}
		return true
	})
}
