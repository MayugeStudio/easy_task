package app

import (
	"bytes"
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
	tests := map[string]struct {
		in      []string
		wantW   string
		wantErr bool
	}{
		"4Tasks": {
			in: []string{
				"- [ ] Task1",
				"- [ ] Task2",
				"- [X] Task3",
				"- [X] Task4",
			},
			wantW: "" +
				"[ ] Task1\n" +
				"[ ] Task2\n" +
				"[X] Task3\n" +
				"[X] Task4\n" +
				"[####################                    ]50.0%",
			wantErr: false,
		},
		"2Tasks1Group": {
			in: []string{
				"- [ ] Task1",
				"- [X] Task2",
				"- GroupTitle",
				"  - [ ] GroupTask1",
				"  - [X] GroupTask2",
			},
			wantW: "" +
				"[ ] Task1\n" +
				"[X] Task2\n" +
				"GroupTitle\n" +
				"  [ ] GroupTask1\n" +
				"  [X] GroupTask2\n" +
				"  [##########          ]50.0%\n" +
				"[####################                    ]50.0%",
			wantErr: false,
		},
		"5Tasks2Group": {
			in: []string{
				"- [X] Task1",
				"- [X] Task2",
				"- [X] Task3",
				"- [ ] Task4",
				"- [ ] Task5",
				"- GroupTitle1",
				"  - [X] GroupTask1",
				"  - [X] GroupTask2",
				"- GroupTitle2",
				"  - [ ] GroupTask1",
				"  - [ ] GroupTask2",
			},
			wantW: "" +
				"[X] Task1\n" +
				"[X] Task2\n" +
				"[X] Task3\n" +
				"[ ] Task4\n" +
				"[ ] Task5\n" +
				"GroupTitle1\n" +
				"  [X] GroupTask1\n" +
				"  [X] GroupTask2\n" +
				"  [####################]100.0%\n" +
				"GroupTitle2\n" +
				"  [ ] GroupTask1\n" +
				"  [ ] GroupTask2\n" +
				"  [                    ]0.0%\n" +
				"[######################                  ]57.1%",
			wantErr: false,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			w := &bytes.Buffer{}
			got := PrintTodoItem(w, "", MockReader{lines: tt.in})
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
	tests := map[string]struct {
		in      []string
		wantErr bool
	}{
		"Tasks": {
			in: []string{
				"- [ ] Task1",
				"- [ ] Task2",
				"- [X] Task3",
				"- [X] Task4",
			},
			wantErr: true,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			w := &ErrorWriter{}
			got := PrintTodoItem(w, "", MockReader{lines: tt.in})
			if (got != nil) != tt.wantErr {
				t.Errorf("PrintTodoItem() got = %v, want %v", got, tt.wantErr)
			}
		})
	}
}
