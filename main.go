package main

import (
	"easy_task/code"
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
	lines, scanErr := code.ScanFile(fileName)
	if scanErr != nil {
		fmt.Printf("scanning file: %v\n", scanErr)
		os.Exit(1)
	}
	tasks := code.ParseStringsToTasks(lines)
	if err := code.PrintTasks(out, tasks); err != nil {
		fmt.Printf("printing tasks: %v\n", err)
		os.Exit(1)
	}
	if err := code.PrintTaskProgress(out, tasks); err != nil {
		fmt.Printf("printing task progress: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Printf("Usage:\n")
	fmt.Printf("\ttst %-12s\t%s", "[filename]", "Displays the task in passed filename.\n")
	fmt.Printf("\ttst %-12s\t%s", "[-h | --help]", "Show this message.\n")
}
