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
	todoList := parse.ToTodoList(lines)
	if err := print.Items(w, todoList.GetItems()); err != nil {
		return fmt.Errorf("printing items: %w", err)
	}
	if err := print.Progress(w, todoList); err != nil {
		return fmt.Errorf("printing progress: %w", err)
	}
	return nil
}
