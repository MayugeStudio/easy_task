package domain

import (
	"reflect"
	"testing"
)

func TestNewTodoItemContainer(t *testing.T) {
	tests := map[string]struct {
		want *TodoItemContainer
	}{
		"Success": {
			&TodoItemContainer{
				make([]*Task, 0),
				make([]*Group, 0),
				0,
			},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewTodoItemContainer()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTodoItemContainer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoItemContainer_AddTask(t *testing.T) {
	type fields struct {
		tasks         []*Task
		groups        []*Group
		doneTaskCount int
	}
	tests := map[string]struct {
		fields fields
		in     *Task
		want   []*Task
	}{
		"Success": {
			fields{
				make([]*Task, 0),
				make([]*Group, 0),
				0,
			},
			&Task{"TaskTitle", false},
			[]*Task{{"TaskTitle", false}},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			c := &TodoItemContainer{
				tt.fields.tasks,
				tt.fields.groups,
				tt.fields.doneTaskCount,
			}
			c.AddTask(tt.in)
			got := c.tasks
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoItemContainer() groups = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoItemContainer_AddGroup(t *testing.T) {
	type fields struct {
		tasks         []*Task
		groups        []*Group
		doneTaskCount int
	}
	tests := map[string]struct {
		fields fields
		in     *Group
		want   []*Group
	}{
		"Success": {
			fields{
				make([]*Task, 0),
				make([]*Group, 0),
				0,
			},
			&Group{"GroupTitle", make([]*Task, 0)},
			[]*Group{{"GroupTitle", make([]*Task, 0)}},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			c := &TodoItemContainer{
				tasks:         tt.fields.tasks,
				groups:        tt.fields.groups,
				doneTaskCount: tt.fields.doneTaskCount,
			}
			c.AddGroup(tt.in)
			got := c.groups
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoItemContainer() groups = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoItemContainer_GetTasks(t *testing.T) {
	type fields struct {
		tasks         []*Task
		groups        []*Group
		doneTaskCount int
	}
	tests := map[string]struct {
		fields fields
		want   []*Task
	}{
		"Success_ZeroTask": {
			fields{
				make([]*Task, 0),
				make([]*Group, 0),
				0,
			},
			[]*Task{},
		},
		"Success_OneTask": {
			fields{
				tasks:         []*Task{{"Task1", false}},
				groups:        make([]*Group, 0),
				doneTaskCount: 0,
			},
			[]*Task{{"Task1", false}},
		},
		"Success_ThreeTasks": {
			fields{
				tasks: []*Task{
					{"Task1", false},
					{"Task2", false},
					{"Task3", true},
				},
				groups:        []*Group{},
				doneTaskCount: 0,
			},
			[]*Task{
				{"Task1", false},
				{"Task2", false},
				{"Task3", true},
			},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			c := &TodoItemContainer{
				tasks:         tt.fields.tasks,
				groups:        tt.fields.groups,
				doneTaskCount: tt.fields.doneTaskCount,
			}
			got := c.GetTasks()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoItemContainer_GetGroups(t *testing.T) {
	type fields struct {
		tasks         []*Task
		groups        []*Group
		doneTaskCount int
	}
	tests := map[string]struct {
		fields fields
		want   []*Group
	}{
		"Success_ZeroGroup": {
			fields{
				make([]*Task, 0),
				make([]*Group, 0),
				0,
			},
			[]*Group{},
		},
		"Success_OneGroup": {
			fields{
				make([]*Task, 0),
				[]*Group{
					{"Group1", make([]*Task, 0)},
				},
				0,
			},
			[]*Group{
				{"Group1", make([]*Task, 0)},
			},
		},
		"Success_ThreeGroups": {
			fields{
				make([]*Task, 0),
				[]*Group{
					{"Group1", make([]*Task, 0)},
					{"Group2", make([]*Task, 0)},
					{"Group3", make([]*Task, 0)},
				},
				0,
			},
			[]*Group{
				{"Group1", make([]*Task, 0)},
				{"Group2", make([]*Task, 0)},
				{"Group3", make([]*Task, 0)},
			},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			c := &TodoItemContainer{
				tasks:         tt.fields.tasks,
				groups:        tt.fields.groups,
				doneTaskCount: tt.fields.doneTaskCount,
			}
			got := c.GetGroups()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
