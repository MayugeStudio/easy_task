package main

import (
	"easy_task/src/app"
	"fmt"
	"os"
)

var out = os.Stdout

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		printUsage()
		os.Exit(1)
	}
	if args[0] == "-h" || args[0] == "--help" {
		printUsage()
		os.Exit(0)
	}
	fileName := args[0]
	outString, exitCode := app.PrintTodoItem(out, fileName)
	fmt.Print(outString)
	os.Exit(exitCode)
}

func printUsage() {
	fmt.Printf("Usage:\n")
	fmt.Printf("\ttst %-12s\t%s", "[filename]", "Displays the task in passed filename.\n")
	fmt.Printf("\ttst %-12s\t%s", "[-h | --help]", "Show this message.\n")
}
