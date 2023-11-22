package app

import (
	"bytes"
	"easy_task/src/logic"
	"fmt"
	"testing"
)

type MockReader struct {
	lines []string
}

func (m MockReader) ReadLines(_ string) (lines []string, err error) {
	return m.lines, nil
}

type ErrorWriter struct{}

func (w ErrorWriter) Write(_ []byte) (n int, err error) {
	return 0, fmt.Errorf("custom error")
}

func TestPrintTodoItem(t *testing.T) {
	type input struct {
		reader logic.FileReader
	}
	tests := map[string]struct {
		in      input
		wantW   string
		wantErr bool
	}{
		"Tasks": {
			input{
				MockReader{
					[]string{
						"- [ ] Task1",
						"- [ ] Task2",
						"- [X] Task3",
						"- [X] Task4",
					},
				},
			},
			"" +
				"[ ] Task1\n" +
				"[ ] Task2\n" +
				"[X] Task3\n" +
				"[X] Task4\n" +
				"[####################                    ]50%",
			false,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			w := &bytes.Buffer{}
			got := PrintTodoItem(w, "", tt.in.reader)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("PrintTodoItem() gotW = %v, want %v", gotW, tt.wantW)
			}
			if (got != nil) != tt.wantErr {
				t.Errorf("PrintTodoItem() got = %v, want %v", got, tt.wantErr)
			}
		})
	}
}

func TestPrintTodoItem_Error(t *testing.T) {
	type input struct {
		reader logic.FileReader
	}
	tests := map[string]struct {
		in      input
		wantErr bool
	}{
		"Tasks": {
			input{
				MockReader{
					[]string{
						"- [ ] Task1",
						"- [ ] Task2",
						"- [X] Task3",
						"- [X] Task4",
					},
				},
			},
			true,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			w := &ErrorWriter{}
			got := PrintTodoItem(w, "", tt.in.reader)
			if (got != nil) != tt.wantErr {
				t.Errorf("PrintTodoItem() got = %v, want %v", got, tt.wantErr)
			}
		})
	}
}
