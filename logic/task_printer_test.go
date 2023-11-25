package logic

import (
	"bytes"
	"github.com/MayugeStudio/easy_task/domain"
	"testing"
)

func TestPrintTasks(t *testing.T) {
	tests := map[string]struct {
		in      []*domain.Task
		wantW   string
		wantErr bool
	}{
		"Success_1Task": {
			in:      []*domain.Task{{"T1", false}},
			wantW:   "[ ] T1\n",
			wantErr: false,
		},
		"Success_2Tasks": {
			in: []*domain.Task{{"T1", false}, {"T2", true}},
			wantW: "" +
				"[ ] T1\n" +
				"[X] T2\n",
			wantErr: false,
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
			gotW := w.String()
			if gotW != tt.wantW {
				t.Errorf("PrintTasks() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_getTaskString(t *testing.T) {
	type input struct {
		task      *domain.Task
		maxLength int
	}
	tests := map[string]struct {
		in   input
		want string
	}{
		"Success_Done":   {in: input{task: &domain.Task{Title: "TaskTitle", IsDone: true}, maxLength: 10}, want: "[X] TaskTitle "},
		"Success_Undone": {in: input{task: &domain.Task{Title: "TaskTitle", IsDone: false}, maxLength: 10}, want: "[ ] TaskTitle "},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := getTaskString(tt.in.task, tt.in.maxLength)
			if got != tt.want {
				t.Errorf("getTaskString() gotW = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMaxTaskNameLength(t *testing.T) {
	tests := map[string]struct {
		in   []*domain.Task
		want int
	}{
		"Success_Length5": {
			in: []*domain.Task{
				{"12", false},
				{"123", false},
				{"12345", false},
			},
			want: 5,
		},
		"Success_Length10": {
			in: []*domain.Task{
				{"1234567890", false},
				{"1234567", false},
				{"123", false},
			},
			want: 10,
		},
		"Success_Length20": {
			in: []*domain.Task{
				{"12345678901234567890", false},
				{"123456789012", false},
				{"123", false},
				{"1234567", false},
			},
			want: 20,
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
