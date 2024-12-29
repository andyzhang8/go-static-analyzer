package main

import (
	"go-static-analyzer/report"
	"go-static-analyzer/rules"
	"go/parser"
	"go/token"
	"log"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Printf("Arguments: %v\n", os.Args)
	fmt.Printf("Number of arguments: %d\n", len(os.Args))
	if len(os.Args) < 2 {
		log.Fatal("Usage: go-static-analyzer <paths-to-go-files-or-directories>")
	}

	fset := token.NewFileSet()

	for _, inputPath := range os.Args[1:] {
		log.Printf("Processing path: %s\n", inputPath)

		fileInfo, err := os.Stat(inputPath)
		if err != nil {
			log.Printf("Failed to access the provided path: %v", err)
			continue
		}

		if fileInfo.IsDir() {
			err := filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if filepath.Ext(path) == ".go" {
					analyzeFile(path, fset)
				}
				return nil
			})
			if err != nil {
				log.Printf("Error analyzing directory %s: %v", inputPath, err)
			}
		} else if filepath.Ext(inputPath) == ".go" {
			analyzeFile(inputPath, fset)
		} else {
			log.Printf("Skipping non-Go file: %s\n", inputPath)
		}
	}

	report.PrintReport()
}

func analyzeFile(filePath string, fset *token.FileSet) {
	log.Printf("Analyzing file: %s\n", filePath)

	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		log.Printf("Failed to parse file %s: %v\n", filePath, err)
		return
	}

	// Apply rules
	rules.CheckComments(node, fset, filePath)
	rules.CheckNestedLoops(node, fset, filePath)
	rules.CheckHardcodedCredentials(node, fset, filePath)
	rules.CheckUnusedVariables(node, fset, filePath)
	rules.CheckFunctionLength(node, fset, filePath)
	rules.CheckErrorHandling(node, fset, filePath)
}
