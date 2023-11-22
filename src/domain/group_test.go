package domain

import (
	"reflect"
	"testing"
)

func TestNewGroup(t *testing.T) {
	tests := map[string]struct {
		in   string
		want *Group
	}{
		"Success": {
			"GroupTitle",
			&Group{Title: "GroupTitle", Tasks: make([]*Task, 0)},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewGroup(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_AddTask(t *testing.T) {
	tests := map[string]struct {
		title string
		in    *Task
		want  []*Task
	}{
		"Success": {
			"GroupTitle",
			&Task{"TaskTitle", false},
			[]*Task{{"TaskTitle", false}},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			g := &Group{
				Title: tt.title,
				Tasks: make([]*Task, 0),
			}

			g.AddTask(tt.in)
			got := g.Tasks
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
