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
		"Success": {in: "GroupTitle", want: &Group{title: "GroupTitle", tasks: make([]*Task, 0)}},
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
		"Success": {title: "G", in: &Task{"T", false}, want: []*Task{{"T", false}}},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			g := &Group{title: tt.title, tasks: make([]*Task, 0)}
			g.AddTask(tt.in)
			got := g.tasks
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_Progress(t *testing.T) {
	tests := map[string]struct {
		isDone []bool
		want   float64
	}{
		"100%":   {isDone: []bool{true, true}, want: 1},
		"50%":    {isDone: []bool{true, false}, want: 0.5},
		"0%":     {isDone: []bool{false, false}, want: 0},
		"NoTask": {isDone: []bool{}, want: 0},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			g := &Group{}
			for _, isDone := range tt.isDone {
				g.tasks = append(g.tasks, &Task{isDone: isDone})
			}
			got := g.Progress()
			if got != tt.want {
				t.Errorf("Progress() = %v, want %v", got, tt.want)
			}
		})
	}
}
