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
			want: &TodoList{tasks: make([]*Task, 0), groups: make([]*Group, 0), items: make([]ProgressItem, 0), doneTaskCount: 0},
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
		items         []ProgressItem
		doneTaskCount int
	}
	tests := map[string]struct {
		fields fields
		in     *Task
		want   []*Task
	}{
		"Success": {
			fields: fields{make([]*Task, 0), make([]*Group, 0), make([]ProgressItem, 0), 0},
			in:     &Task{"T", false}, want: []*Task{{"T", false}},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			c := &TodoList{tt.fields.tasks, tt.fields.groups, tt.fields.items, tt.fields.doneTaskCount}
			c.AddTask(tt.in)
			got := c.tasks
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoList() tasks = %v, want %v", got, tt.want)
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
			in:     &Group{"G", make([]*Task, 0)}, want: []*Group{{"G", make([]*Task, 0)}},
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
			fields: fields{tasks: []*Task{{"T1", false}}, groups: make([]*Group, 0), doneTaskCount: 0},
			want:   []*Task{{"T1", false}},
		},
		"Success_ThreeTasks": {
			fields: fields{tasks: []*Task{{title: "T1", isDone: false}, {title: "T2", isDone: false}, {title: "T3", isDone: true}}, groups: []*Group{}, doneTaskCount: 0},
			want:   []*Task{{title: "T1", isDone: false}, {title: "T2", isDone: false}, {title: "T3", isDone: true}},
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
			fields: fields{make([]*Task, 0), []*Group{{"G1", make([]*Task, 0)}}, 0},
			want:   []*Group{{"G1", make([]*Task, 0)}},
		},
		"Success_ThreeGroups": {
			fields: fields{make([]*Task, 0), []*Group{{"G1", make([]*Task, 0)}, {"G2", make([]*Task, 0)}, {"G3", make([]*Task, 0)}}, 0},
			want:   []*Group{{"G1", make([]*Task, 0)}, {"G2", make([]*Task, 0)}, {"G3", make([]*Task, 0)}},
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

func TestTodoList_Progress(t *testing.T) {
	type input struct {
		tasks      []bool
		groupTasks [][]bool
	}
	tests := map[string]struct {
		in   input
		want float64
	}{
		"100%_TaskOnly": {
			in: input{
				tasks: []bool{true, true},
			},
			want: 1,
		},
		"100%_GroupOnly": {
			in: input{
				groupTasks: [][]bool{{true, true}, {true, true}},
			},
			want: 1,
		},
		"100%_Mix": {
			in: input{
				tasks:      []bool{true, true},
				groupTasks: [][]bool{{true, true}, {true, true}},
			},
			want: 1,
		},
		"50%_TaskOnly": {
			in: input{
				tasks: []bool{true, true, false, false},
			},
			want: 0.5,
		},
		"50%_GroupOnly_0011": {
			in: input{
				groupTasks: [][]bool{{false, false}, {true, true}},
			},
			want: 0.5,
		},
		"50%_GroupOnly_0101": {
			in: input{
				groupTasks: [][]bool{{false, true}, {false, true}},
			},
			want: 0.5,
		},
		"50%_Mix_010101": {
			in: input{
				tasks:      []bool{false, true},
				groupTasks: [][]bool{{false, true}, {false, true}},
			},
			want: 0.5,
		},
		"0%_TaskOnly": {
			in: input{
				tasks: []bool{false, false},
			},
			want: 0,
		},
		"0%_GroupOnly": {
			in: input{
				groupTasks: [][]bool{{false, false}, {false, false}},
			},
			want: 0,
		},
		"0%_Mix": {
			in: input{
				tasks:      []bool{false},
				groupTasks: [][]bool{{false, false}, {false, false}},
			},
			want: 0,
		},
		"0%_NoItem": {
			in: input{
				tasks:      nil,
				groupTasks: nil,
			},
			want: 0,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			c := &TodoList{}
			for _, status := range tt.in.tasks {
				task := &Task{isDone: status}
				c.AddTask(task)
			}
			for _, areTasksDone := range tt.in.groupTasks {
				group := &Group{}
				for _, isDone := range areTasksDone {
					group.AddTask(&Task{isDone: isDone})
				}
				c.AddGroup(group)
			}

			got := c.Progress()

			if got != tt.want {
				t.Errorf("Progress() = %v, want %v", got, tt.want)
			}
		})
	}
}
