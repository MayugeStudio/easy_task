package code

import (
	"reflect"
	"testing"
)

func TestParseStringsToTasks(t *testing.T) {
	lines := []string{"- [ ] Task1", "- [ ] Task2", "- [X] Task3"}
	expectedTasks := []*Task{
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

	actualTasks := ParseStringsToTasks(lines)

	if !reflect.DeepEqual(actualTasks, expectedTasks) {
		t.Errorf("ParseLines() got = %v, want %v", actualTasks, expectedTasks)
	}
}

func TestParseStringToTask(t *testing.T) {
	tokens := []string{"X", "TaskName"}
	expectedTask := &Task{Title: "TaskName", IsDone: true}
	if actualTask := ParseStringToTask(tokens); !reflect.DeepEqual(actualTask, expectedTask) {
		t.Errorf("parseLine() = %v, want %v", actualTask, expectedTask)
	}
}
