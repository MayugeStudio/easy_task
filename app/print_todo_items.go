package app

import (
	logic2 "easy_task/logic"
	"fmt"
	"io"
)

func PrintTodoItem(w io.Writer, fileName string, reader logic2.FileReader) error {
	lines, scanErr := logic2.ScanFile(fileName, reader)
	if scanErr != nil {
		return fmt.Errorf("scanning file: %w", scanErr)
	}
	todoItemContainer := logic2.ParseStringsToTasks(lines)
	if err := logic2.PrintTasks(w, todoItemContainer.GetTasks()); err != nil {
		return fmt.Errorf("printing tasks: %w", err)
	}
	if err := logic2.PrintGroups(w, todoItemContainer.GetGroups()); err != nil {
		return fmt.Errorf("printing groups: %w", err)
	}
	if err := logic2.PrintProgress(w, todoItemContainer); err != nil {
		return fmt.Errorf("printing progress: %w", err)
	}
	return nil
}
