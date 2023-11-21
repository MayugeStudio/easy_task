package main

import (
	"easy_task/src/logic"
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
	lines, scanErr := logic.ScanFile(fileName, logic.FileScanner{})
	if scanErr != nil {
		fmt.Printf("scanning file: %v\n", scanErr)
		os.Exit(1)
	}
	todoItemContainer := logic.ParseStringsToTasks(lines)
	if err := logic.PrintTasks(out, todoItemContainer.GetTasks()); err != nil {
		fmt.Printf("printing tasks: %v\n", err)
		os.Exit(1)
	}
	if err := logic.PrintTaskProgress(out, todoItemContainer.GetTasks()); err != nil {
		fmt.Printf("printing task progress: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Printf("Usage:\n")
	fmt.Printf("\ttst %-12s\t%s", "[filename]", "Displays the task in passed filename.\n")
	fmt.Printf("\ttst %-12s\t%s", "[-h | --help]", "Show this message.\n")
}
