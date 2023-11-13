package code

import "testing"

func TestPrintTaskProgress(t *testing.T) {
	tasks := []TaskPtr{{IsDone: false}, {IsDone: false}, {IsDone: true}}
	expected := "[#############                           ]33%"
	testWriter := &TestWriter{}
	if err := PrintTaskProgress(testWriter, tasks); err != nil {
		t.Fatalf("PrintTaskProgress() error = %v", err)
	}
	actual := string(testWriter.WrittenData)
	if expected != actual {
		t.Errorf("PrintTaskProgress() = %s, want %s", actual, expected)
	}
}
