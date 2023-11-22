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
