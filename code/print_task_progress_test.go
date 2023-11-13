package code

import "testing"

func TestPrintTaskProgress(t *testing.T) {
	tasks := []TaskPtr{{IsDone: false}, {IsDone: false}, {IsDone: true}}
	expected := "[#############                           ]33%"
	writer := &TestWriter{}
	if err := PrintTaskProgress(writer, tasks); err != nil {
		t.Fatalf("PrintTaskProgress() error = %v", err)
	}
	actual := string(writer.WrittenData)
	if expected != actual {
		t.Errorf("PrintTaskProgress() = %s, want %s", actual, expected)
	}
}
