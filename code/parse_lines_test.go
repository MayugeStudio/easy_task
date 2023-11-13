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
	expectedErrorMessages := make([]string, 0)

	actualTasks, actualErrorMessages := ParseLines(lines)

	if !reflect.DeepEqual(actualTasks, expectedTasks) {
		t.Errorf("ParseLines() got = %v, want %v", actualTasks, expectedTasks)
	}
	if !reflect.DeepEqual(actualErrorMessages, expectedErrorMessages) {
		t.Errorf("ParseLines() got1 = %v, want %v", actualErrorMessages, expectedErrorMessages)
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
	actualTokens, err := processedLineToTokens(processedLine)
	if err != nil {
		t.Errorf("processedLineToTokens() error = %v", err)
		return
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
