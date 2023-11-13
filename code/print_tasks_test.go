package code

import (
	"testing"
)

func TestPrintTasks(t *testing.T) {
	tasks := []TaskPtr{{Title: "TestTask1", IsDone: true}, {Title: "TestTask2", IsDone: false}}
	writer := &TestWriter{}
	expected := "[X] TestTask1\n[ ] TestTask2\n"
	err := PrintTasks(writer, tasks)
	if err != nil {
		t.Fatalf("PrintTasks() error = %v", err)
	}
	if actual := string(writer.WrittenData); actual != expected {
		t.Errorf("PrintTasks() = %v, want %v", actual, expected)
	}
}

func Test_printTask(t *testing.T) {
	task := &Task{Title: "Test", IsDone: true}
	writer := &TestWriter{}
	expected := "[X] Test\n"
	err := printTask(writer, task, len(task.Title))
	if err != nil {
		t.Fatalf("printTask() error = %v", err)
	}
	if actual := string(writer.WrittenData); actual != expected {
		t.Errorf("printTask() = %v, want %v", actual, expected)
	}
}

func Test_getMaxTaskNameLength(t *testing.T) {
	tasks := []TaskPtr{{Title: "10Length--"}, {Title: "8Length-"}, {Title: "7Length"}, {Title: "4Len"}}
	expected := 10
	if actual := getMaxTaskNameLength(tasks); actual != expected {
		t.Errorf("getMaxTaskNameLength() = %v, want %v", actual, expected)
	}
}
