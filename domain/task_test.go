package domain

import (
	"reflect"
	"testing"
)

func TestNewTask(t *testing.T) {
	t.Parallel()
	type args struct {
		title  string
		isDone bool
	}
	tests := map[string]struct {
		in   args
		want *Task
	}{
		"Success": {in: args{title: "T", isDone: false}, want: &Task{title: "T", isDone: false}},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewTask(tt.in.title, tt.in.isDone)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_Progress(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		isDone bool
		want   float64
	}{
		"100%": {isDone: true, want: 1},
		"0%":   {isDone: false, want: 0},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			task := &Task{isDone: tt.isDone}
			if got := task.Progress(); got != tt.want {
				t.Errorf("Progress() = %v, want %v", got, tt.want)
			}
		})
	}
}
