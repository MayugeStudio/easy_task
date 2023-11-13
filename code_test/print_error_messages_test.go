package code_test

import (
	"easy_task/code"
	"testing"
)

type TestWriter struct {
	WrittenData []byte
}

func (tw *TestWriter) Write(p []byte) (n int, err error) {
	tw.WrittenData = append(tw.WrittenData, p...)
	return len(p), nil
}

func TestPrintErrorMessages(t *testing.T) {
	type args struct {
		messages []string
	}
	testCases := []struct {
		testName string
		args     args
	}{
		{
			testName: "Ok",
			args:     args{messages: []string{"Error1", "Error2"}},
		},
	}
	testWriter := &TestWriter{}
	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			if err := code.PrintErrorMessages(testWriter, tt.args.messages); err != nil {
				t.Fatalf("PrintErrorMessages error = %v", err)
			}
		})
	}
}
