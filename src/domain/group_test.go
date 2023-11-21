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

func TestNewGroupContainer(t *testing.T) {
	tests := map[string]struct {
		want *GroupContainer
	}{
		"Success": {
			&GroupContainer{
				groups: make([]*Group, 0),
			},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewGroupContainer()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGroupContainer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupContainer_AddGroup(t *testing.T) {
	tests := map[string]struct {
		groups []*Group
		in     *Group
		want   []*Group
	}{
		"Success": {
			[]*Group{},
			&Group{Title: "NewGroup"},
			[]*Group{{Title: "NewGroup"}},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			c := &GroupContainer{
				groups: tt.groups,
			}
			c.AddGroup(tt.in)
			got := c.groups
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGroupContainer() groups = %v, want %v", got, tt.want)
			}
		})
	}
}
