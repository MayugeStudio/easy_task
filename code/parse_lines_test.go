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
	if actualTask := parseLine(tokens); !reflect.DeepEqual(actualTask, expectedTask) {
		t.Errorf("parseLine() = %v, want %v", actualTask, expectedTask)
	}
}

func Test_processedLineToTokens_success(t *testing.T) {
	processedLine := " X TestName"
	expectedTokens := []string{"X", "TestName"}
	actualTokens, skip := processedLineToTokens(processedLine)
	if skip == true {
		t.Errorf("processedLineToTokens() skip = true, want false")
	}
	if !reflect.DeepEqual(actualTokens, expectedTokens) {
		t.Errorf("processedLineToTokens() actual = %v, want %v", actualTokens, expectedTokens)
	}
}

func Test_removeUnnecessaryTokenFromLine_success(t *testing.T) {
	line := "- [X] TaskName"
	expectedProcessedLine := " X TaskName"
	actualProcessedLine, isSkip := removeUnnecessaryTokenFromLine(line)
	if actualProcessedLine != expectedProcessedLine {
		t.Errorf("removeUnnecessaryTokenFromLine() actualProcessedLine = %v, want %v", actualProcessedLine, expectedProcessedLine)
	}
	if isSkip != false {
		t.Errorf("removeUnnecessaryTokenFromLine() IsSkip = %v, want %v", isSkip, false)
	}
}
