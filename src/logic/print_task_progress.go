package logic

import (
	"easy_task/src/domain"
	"fmt"
	"io"
	"strings"
)

const progressSymbol = "#"
const defaultProgressBarLength = 40.0

func PrintTaskProgress(w io.Writer, tasks []*domain.Task) error {
	if len(tasks) == 0 {
		return nil
	}
	taskProgressString := getTaskProgressString(tasks, defaultProgressBarLength)
	if _, err := fmt.Fprint(w, taskProgressString); err != nil {
		return err
	}
	return nil
}

func getTaskProgressString(tasks []*domain.Task, length float64) string {
	taskNum := float64(len(tasks))
	doneTaskNum := 0.0
	for _, task := range tasks {
		if task.IsDone {
			doneTaskNum++
		}
	}
	doneTaskRatio := doneTaskNum / taskNum
	doneTaskStrLength := int(doneTaskRatio * length)
	doneTaskStr := strings.Repeat(progressSymbol, doneTaskStrLength)
	undoneTaskStr := strings.Repeat(" ", int(length)-doneTaskStrLength)
	return fmt.Sprintf("[%s%s]%d%%", doneTaskStr, undoneTaskStr, int(doneTaskRatio*100))
}
