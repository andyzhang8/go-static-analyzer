package main

// Test function
func testFunction() {
	// Nested loop example
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			println(i, j)
		}
	}

	// Dummy call to simulate error handling
	_ = someFunction()
}

// Dummy function to simulate error returning
func someFunction() error {
	return nil
}
