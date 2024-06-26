package main

import (
	"fmt"
	"os"
)

var (
	programName    = "mock"
	programVersion = "v3.1"
	// ExecutionTimestamp provides the current time to generate log files
	ExecutionTimestamp = TimeNow()
	// Path set the path or the location where the files will be generated
	Path = fmt.Sprintf("%s/%s/%s", os.Getenv("HOME"), programName, ExecutionTimestamp)
)

// The main function block
func main() {
	// Execute the cobra CLI & run the program
	err := rootCmd.Execute()
	if err != nil {
		Fatalf("Error executing the mock data command cli, err: %v", err)
	}
}
