package main

import (
	"easy_task/code"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Printf("Usage: tst [filename]")
		os.Exit(1)
	}
	fileName := args[0]
	lines, scanErr := code.ScanFile(fileName)
	if scanErr != nil {
		fmt.Println("Error:", scanErr)
		os.Exit(1)
	}
	tasks, msgSlice := code.ParseLines(lines)
	if err := code.PrintErrorMessages(os.Stdout, msgSlice); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if err := code.PrintTasks(os.Stdout, tasks); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if err := code.PrintTaskProgress(os.Stdout, tasks); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
