package app

import (
	"fmt"
	"io"

	"github.com/MayugeStudio/easy_task/logic"
)

func PrintTodoItem(w io.Writer, fileName string, reader logic.FileReader) error {
	lines, scanErr := logic.ScanFile(fileName, reader)
	if scanErr != nil {
		return fmt.Errorf("scanning file: %w", scanErr)
	}
	todoItemContainer := logic.ParseStringsToTasks(lines)
	if err := logic.PrintTasks(w, todoItemContainer.GetTasks()); err != nil {
		return fmt.Errorf("printing tasks: %w", err)
	}
	if err := logic.PrintGroups(w, todoItemContainer.GetGroups()); err != nil {
		return fmt.Errorf("printing groups: %w", err)
	}
	if err := logic.PrintProgress(w, todoItemContainer); err != nil {
		return fmt.Errorf("printing progress: %w", err)
	}
	return nil
}
