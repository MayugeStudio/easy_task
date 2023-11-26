package app

import (
	"fmt"
	print2 "github.com/MayugeStudio/easy_task/logic/print"
	"io"

	"github.com/MayugeStudio/easy_task/logic"
	"github.com/MayugeStudio/easy_task/logic/parse"
)

func PrintTodoItem(w io.Writer, fileName string, reader logic.FileReader) error {
	lines, scanErr := logic.ScanFile(fileName, reader)
	if scanErr != nil {
		return fmt.Errorf("scanning file: %w", scanErr)
	}
	todoItemContainer := parse.ToTodoList(lines)
	if err := print2.Tasks(w, todoItemContainer.GetTasks()); err != nil {
		return fmt.Errorf("printing tasks: %w", err)
	}
	if err := print2.Groups(w, todoItemContainer.GetGroups()); err != nil {
		return fmt.Errorf("printing groups: %w", err)
	}
	if err := print2.Progress(w, todoItemContainer); err != nil {
		return fmt.Errorf("printing progress: %w", err)
	}
	return nil
}
