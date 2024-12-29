package main

import "fmt"

func main() {
	err := someFunction()
	fmt.Println("This is unrelated")

	if err != nil {
		fmt.Println("Error occurred:", err) // Proper error handling
	}

	anotherErr := someFunction()
	fmt.Println("Ignoring error:", anotherErr) // Improper error handling
}

func someFunction() error {
	return fmt.Errorf("an error occurred")
}
