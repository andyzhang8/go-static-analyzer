package rules

import (
	"go-static-analyzer/report"
	"go-static-analyzer/utils"
	"go/ast"
	"go/token"
)

// Flags unused variables
func CheckUnusedVariables(node *ast.File, fset *token.FileSet, filePath string) {
	ast.Inspect(node, func(n ast.Node) bool {
		if decl, ok := n.(*ast.GenDecl); ok {
			for _, spec := range decl.Specs {
				if valueSpec, ok := spec.(*ast.ValueSpec); ok {
					for _, name := range valueSpec.Names {
						if !name.IsExported() {
							report.AddIssue(filePath, "Unused Variable", utils.GetLine(name.Pos(), fset), 0, "Variable '"+name.Name+"' is declared but not used")
						}
					}
				}
			}
		}
		return true
	})
}
