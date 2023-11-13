package code

import (
	"fmt"
	"io"
	"strings"
)

// PrintTaskProgress prints a progress bar to represent the completion status of tasks.
// It takes a slice of TaskPtr and calculates the ratio of completed tasks to total tasks.
func PrintTaskProgress(w io.Writer, tasks []TaskPtr) error {
	progressBarLength := DefaultProgressBarLength
	taskNum := float64(len(tasks))
	doneTaskNum := 0.0
	for _, task := range tasks {
		if task.IsDone {
			doneTaskNum++
		}
	}
	doneTaskRatio := doneTaskNum / taskNum
	doneTaskStrLength := int(doneTaskRatio * progressBarLength)
	doneTaskStr := strings.Repeat(ProgressSymbol, doneTaskStrLength)
	undoneTaskStr := strings.Repeat(" ", int(progressBarLength)-doneTaskStrLength)
	if _, err := fmt.Fprintf(w, "[%s%s]%d%%", doneTaskStr, undoneTaskStr, int(doneTaskRatio*100)); err != nil {
		return err
	}
	return nil
}
