package code

import (
	"reflect"
	"testing"
)

func TestParseLines(t *testing.T) {
	lines := []string{"- [ ] Task1", "- [ ] Task2", "- [X] Task3"}
	expectedTasks := []TaskPtr{
		{
			Title:  "Task1",
			IsDone: false,
		},
		{
			Title:  "Task2",
			IsDone: false,
		},
		{
			Title:  "Task3",
			IsDone: true,
		},
	}

	actualTasks := ParseLines(lines)

	if !reflect.DeepEqual(actualTasks, expectedTasks) {
		t.Errorf("ParseLines() got = %v, want %v", actualTasks, expectedTasks)
	}
}

func Test_parseLine(t *testing.T) {
	tokens := []string{"X", "TaskName"}
	expectedTask := TaskPtr(&Task{Title: "TaskName", IsDone: true})
	if actualTask := convertTokensToTask(tokens); !reflect.DeepEqual(actualTask, expectedTask) {
		t.Errorf("convertTokensToTask() = %v, want %v", actualTask, expectedTask)
	}
}

func Test_removeUnnecessaryTokenFromLine_success(t *testing.T) {
	line := "- [X] TaskName"
	expectedProcessedLine := " X TaskName"
	actualProcessedLine := processLine(line)
	if actualProcessedLine != expectedProcessedLine {
		t.Errorf("processLine() actualProcessedLine = %v, want %v", actualProcessedLine, expectedProcessedLine)
	}
}
