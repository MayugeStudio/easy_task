package domain

import (
	"reflect"
	"testing"
)

func TestNewTask(t *testing.T) {
	type args struct {
		title  string
		isDone bool
	}
	tests := map[string]struct {
		in   args
		want *Task
	}{
		"Success": {in: args{"TaskTitle", false}, want: &Task{"TaskTitle", false}},
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

func TestTask_Progress(t1 *testing.T) {
	tests := map[string]struct {
		isDone bool
		want   float64
	}{
		"100%": {isDone: true, want: 1},
		"0%":   {isDone: false, want: 0},
	}
	for testName, tt := range tests {
		t1.Run(testName, func(t1 *testing.T) {
			t := &Task{IsDone: tt.isDone}
			if got := t.Progress(); got != tt.want {
				t1.Errorf("Progress() = %v, want %v", got, tt.want)
			}
		})
	}
}
