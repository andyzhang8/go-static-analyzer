package utils

import (
	"go/token"
	"strconv"
)

// Extracts the line number from a token.Pos
func GetLine(pos token.Pos, fset *token.FileSet) int {
	return fset.Position(pos).Line
}

// Extracts the column number from a token.Pos
func GetColumn(pos token.Pos, fset *token.FileSet) int {
	return fset.Position(pos).Column
}

// Removes surrounding quotes from a string
func TrimQuotes(s string) string {
	if len(s) > 1 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}

// Converts an integer to a string
func IntToString(num int) string {
	return strconv.Itoa(num)
}
