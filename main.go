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
	lines := code.ScanFile(fileName)
	tasks, msgSlice := code.ParseLines(lines)
	if err := code.PrintErrorMessages(os.Stdout, msgSlice); err != nil {
		fmt.Println("Error:", err)
	}
	code.PrintTasks(tasks)
	if err := code.PrintTaskProgress(os.Stdout, tasks); err != nil {
		fmt.Println("Error:", err)
	}
}
