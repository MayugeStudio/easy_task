package main

import (
	"easy_task/code"
	"fmt"
	"os"
)

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
		fmt.Println("Error:", scanErr)
		os.Exit(1)
	}
	tasks := code.ParseTaskStringsToTasks(lines)
	if err := code.PrintTasks(os.Stdout, tasks); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if err := code.PrintTaskProgress(os.Stdout, tasks); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Printf("Usage:\n")
	fmt.Printf("\ttst %-12s\t%s", "[filename]", "Displays the task in passed filename.\n")
	fmt.Printf("\ttst %-12s\t%s", "[-h | --help]", "Show this message.\n")
}
