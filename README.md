# Silent Errors with Missing Map Keys in Go

This repository demonstrates a common, yet subtle, error in Go: accessing a map with a non-existent key.  Go's map access silently returns the zero value for the corresponding data type if the key is not found. This behavior can lead to hard-to-debug errors, as the program doesn't crash, but produces unexpected results.

The `bug.go` file shows the problematic code, which accesses a map without checking for the presence of the key.  `bugSolution.go` provides a corrected version that uses the comma ok idiom to handle missing keys gracefully.

## How to reproduce the bug

1. Clone the repository.
2. Navigate to the project directory.
3. Run the following command: `go run bug.go`
4. Observe the output which is 0 but one may expect an error.

## How to fix the bug

Refer to `bugSolution.go` for a solution using the comma ok idiom.