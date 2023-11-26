package app

import (
	"fmt"
	"io"

	"github.com/MayugeStudio/easy_task/logic"
	"github.com/MayugeStudio/easy_task/logic/parse"
	"github.com/MayugeStudio/easy_task/logic/print"
)

func PrintTodoItem(w io.Writer, fileName string, reader logic.FileReader) error {
	lines, scanErr := logic.ScanFile(fileName, reader)
	if scanErr != nil {
		return fmt.Errorf("scanning file: %w", scanErr)
	}
	todoItemContainer := parse.ToTodoList(lines)
	if err := print.Tasks(w, todoItemContainer.GetTasks()); err != nil {
		return fmt.Errorf("printing tasks: %w", err)
	}
	if err := print.Groups(w, todoItemContainer.GetGroups()); err != nil {
		return fmt.Errorf("printing groups: %w", err)
	}
	if err := print.Progress(w, todoItemContainer); err != nil {
		return fmt.Errorf("printing progress: %w", err)
	}
	return nil
}
