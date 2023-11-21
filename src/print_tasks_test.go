package src

import (
	"bytes"
	"testing"
)

func TestPrintTasks(t *testing.T) {
	tests := map[string]struct {
		in      []*Task
		wantW   string
		wantErr bool
	}{
		"Success_1Task": {
			[]*Task{
				{"Task1", false},
			},
			"[ ] Task1\n",
			false,
		},
		"Success_3Tasks": {
			[]*Task{
				{"Task1", false},
				{"Task2", true},
				{"Task3", true},
			},
			"[ ] Task1\n[X] Task2\n[X] Task3\n",
			false,
		},
		"Success_10Tasks": {
			[]*Task{
				{"0Hi", false},
				{"1BuyTheMilk", false},
				{"2MaxLengthName", false},
				{"3ILikeSinging", true},
				{"4I'm Gopher", true},
			},
			"" +
				"[ ] 0Hi           \n" +
				"[ ] 1BuyTheMilk   \n" +
				"[ ] 2MaxLengthName\n" +
				"[X] 3ILikeSinging \n" +
				"[X] 4I'm Gopher   \n",
			false,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			w := &bytes.Buffer{}
			err := PrintTasks(w, tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrintTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("PrintTasks() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_printTask(t *testing.T) {
	type input struct {
		task              *Task
		maxTaskNameLength int
	}
	tests := map[string]struct {
		in      input
		wantW   string
		wantErr bool
	}{
		"Success_Done": {
			input{&Task{"TaskTitle", true}, 10},
			"[X] TaskTitle \n",
			false,
		},
		"Success_Undone": {
			input{&Task{"TaskTitle", false}, 10},
			"[ ] TaskTitle \n",
			false,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			w := &bytes.Buffer{}
			err := printTask(w, tt.in.task, tt.in.maxTaskNameLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("printTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("printTask() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_getMaxTaskNameLength(t *testing.T) {
	tests := map[string]struct {
		in   []*Task
		want int
	}{
		"Success_Length5": {
			[]*Task{
				{"12", false},
				{"123", false},
				{"12345", false},
			},
			5,
		},
		"Success_Length10": {
			[]*Task{
				{"1234567890", false},
				{"1234567", false},
				{"123", false},
			},
			10,
		},
		"Success_Length20": {
			[]*Task{
				{"12345678901234567890", false},
				{"123456789012", false},
				{"123", false},
				{"1234567", false},
			},
			20,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := getMaxTaskNameLength(tt.in); got != tt.want {
				t.Errorf("getMaxTaskNameLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
