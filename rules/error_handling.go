package rules

import (
	"go-static-analyzer/report"
	"go-static-analyzer/utils"
	"go/ast"
	"go/token"
)

// Ensures errors are handled properly
func CheckErrorHandling(node *ast.File, fset *token.FileSet, filePath string) {
	ast.Inspect(node, func(n ast.Node) bool {
		if assign, ok := n.(*ast.AssignStmt); ok {
			for _, lhs := range assign.Lhs {
				if ident, ok := lhs.(*ast.Ident); ok && ident.Name == "err" {
					if !isErrorHandled(node, assign, ident.Name) {
						report.AddIssue(filePath, "Improper Error Handling", utils.GetLine(assign.Pos(), fset), 0, "Error returned is not handled")
					}
				}
			}
		}
		return true
	})
}

// Checks if an error variable is used in subsequent statements
func isErrorHandled(funcNode ast.Node, assign *ast.AssignStmt, errVar string) bool {
	block, ok := findEnclosingBlock(funcNode, assign)
	if !ok {
		return false
	}

	foundAssign := false
	for _, stmt := range block.List {
		if stmt == assign {
			foundAssign = true
			continue
		}

		if foundAssign {
			if usesVariable(stmt, errVar) {
				return true
			}
		}
	}

	return false
}

// Finds the enclosing block for the given assignment
func findEnclosingBlock(node ast.Node, target ast.Node) (*ast.BlockStmt, bool) {
	var enclosingBlock *ast.BlockStmt
	ast.Inspect(node, func(n ast.Node) bool {
		if block, ok := n.(*ast.BlockStmt); ok {
			for _, stmt := range block.List {
				if stmt == target {
					enclosingBlock = block
					return false
				}
			}
		}
		return true
	})
	return enclosingBlock, enclosingBlock != nil
}

// Checks if a statement uses a specific variable
func usesVariable(stmt ast.Node, varName string) bool {
	found := false
	ast.Inspect(stmt, func(n ast.Node) bool {
		if ident, ok := n.(*ast.Ident); ok && ident.Name == varName {
			found = true
			return false
		}
		return true
	})
	return found
}
