package logic

import (
	"easy_task/src/domain"
	"fmt"
	"io"
	"strings"
)

const ProgressSymbol = "#"
const DefaultProgressBarLength = 40.0

func PrintTaskProgress(w io.Writer, tasks []*domain.Task) error {
	if len(tasks) == 0 {
		return nil
	}
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