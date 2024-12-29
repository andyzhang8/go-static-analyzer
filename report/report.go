package report

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// Issue struct
type Issue struct {
	Rule   string
	Line   int
	Column int
	Detail string
}

// Maps file names to their respective issues
var FileReport = make(map[string][]Issue)

// Adds a new issue to the list for a specific file
func AddIssue(fileName, rule string, line, column int, detail string) {
	FileReport[fileName] = append(FileReport[fileName], Issue{
		Rule:   rule,
		Line:   line,
		Column: column,
		Detail: detail,
	})
}

// Generates and prints the report for all files
func PrintReport() {
	if len(FileReport) == 0 {
		fmt.Println("No issues detected. Great job!")
		return
	}

	fmt.Println("\nStatic Code Analysis Report:")
	fmt.Println("============================")

	for file, issues := range FileReport {
		fmt.Printf("\nFile: %s\n", file)
		writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)
		fmt.Fprintln(writer, "Rule\tLine\tColumn\tDetail")

		for _, issue := range issues {
			fmt.Fprintf(writer, "%s\t%d\t%d\t%s\n", issue.Rule, issue.Line, issue.Column, issue.Detail)
		}

		writer.Flush()
	}
}
