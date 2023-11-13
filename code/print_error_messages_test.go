package code

import (
	"testing"
)

func TestPrintErrorMessages(t *testing.T) {
	messages := []string{"Error1", "Error2"}
	writer := &TestWriter{}
	if err := PrintErrorMessages(writer, messages); err != nil {
		t.Fatalf("PrintErrorMessages error = %v", err)
	}
}
