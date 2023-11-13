package code

import (
	"testing"
)

func TestPrintErrorMessages(t *testing.T) {
	messages := []string{"Error1", "Error2"}
	testWriter := &TestWriter{}
	if err := PrintErrorMessages(testWriter, messages); err != nil {
		t.Fatalf("PrintErrorMessages error = %v", err)
	}
}
