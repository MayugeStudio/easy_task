package domain

import (
	"reflect"
	"testing"
)

func TestNewTodoList(t *testing.T) {
	tests := map[string]struct {
		want *TodoList
	}{
		"Success": {
			want: &TodoList{tasks: make([]*Task, 0), groups: make([]*Group, 0), doneTaskCount: 0},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewTodoList()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTodoList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoList_AddTask(t *testing.T) {
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
			fields: fields{make([]*Task, 0), make([]*Group, 0), 0},
			in:     &Task{"TaskTitle", false},
			want:   []*Task{{"TaskTitle", false}},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			c := &TodoList{tt.fields.tasks, tt.fields.groups, tt.fields.doneTaskCount}
			c.AddTask(tt.in)
			got := c.tasks
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoList() groups = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoList_AddGroup(t *testing.T) {
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
			fields: fields{make([]*Task, 0), make([]*Group, 0), 0},
			in:     &Group{"GroupTitle", make([]*Task, 0)},
			want:   []*Group{{"GroupTitle", make([]*Task, 0)}},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			c := &TodoList{tasks: tt.fields.tasks, groups: tt.fields.groups, doneTaskCount: tt.fields.doneTaskCount}
			c.AddGroup(tt.in)
			got := c.groups
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoList() groups = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoList_GetTasks(t *testing.T) {
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
			fields: fields{make([]*Task, 0), make([]*Group, 0), 0},
			want:   []*Task{},
		},
		"Success_OneTask": {
			fields: fields{tasks: []*Task{{"Task1", false}}, groups: make([]*Group, 0), doneTaskCount: 0},
			want:   []*Task{{"Task1", false}},
		},
		"Success_ThreeTasks": {
			fields: fields{tasks: []*Task{{"Task1", false}, {"Task2", false}, {"Task3", true}}, groups: []*Group{}, doneTaskCount: 0},
			want:   []*Task{{"Task1", false}, {"Task2", false}, {"Task3", true}},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			c := &TodoList{tasks: tt.fields.tasks, groups: tt.fields.groups, doneTaskCount: tt.fields.doneTaskCount}
			got := c.GetTasks()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoList_GetGroups(t *testing.T) {
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
			fields: fields{make([]*Task, 0), make([]*Group, 0), 0},
			want:   []*Group{},
		},
		"Success_OneGroup": {
			fields: fields{make([]*Task, 0), []*Group{{"Group1", make([]*Task, 0)}}, 0},
			want:   []*Group{{"Group1", make([]*Task, 0)}},
		},
		"Success_ThreeGroups": {
			fields: fields{make([]*Task, 0), []*Group{{"Group1", make([]*Task, 0)}, {"Group2", make([]*Task, 0)}, {"Group3", make([]*Task, 0)}}, 0},
			want:   []*Group{{"Group1", make([]*Task, 0)}, {"Group2", make([]*Task, 0)}, {"Group3", make([]*Task, 0)}},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			c := &TodoList{tasks: tt.fields.tasks, groups: tt.fields.groups, doneTaskCount: tt.fields.doneTaskCount}
			got := c.GetGroups()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
