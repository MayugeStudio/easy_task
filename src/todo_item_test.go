package src

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
				NewTaskContainer(),
				NewGroupContainer(),
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
		taskContainer  *TaskContainer
		groupContainer *GroupContainer
	}
	tests := map[string]struct {
		fields fields
		in     *Task
		want   []*Task
	}{
		"Success": {
			fields{NewTaskContainer(), NewGroupContainer()},
			&Task{"TaskTitle", false},
			[]*Task{{"TaskTitle", false}},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			c := &TodoItemContainer{
				taskContainer:  tt.fields.taskContainer,
				groupContainer: tt.fields.groupContainer,
			}
			c.AddTask(tt.in)
			got := c.taskContainer.tasks
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoItemContainer() groups = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoItemContainer_AddGroup(t *testing.T) {
	type fields struct {
		taskContainer  *TaskContainer
		groupContainer *GroupContainer
	}
	tests := map[string]struct {
		fields fields
		in     *Group
		want   []*Group
	}{
		"Success": {
			fields{NewTaskContainer(), NewGroupContainer()},
			&Group{"GroupTitle", make([]*Task, 0)},
			[]*Group{{"GroupTitle", make([]*Task, 0)}},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			c := &TodoItemContainer{
				taskContainer:  tt.fields.taskContainer,
				groupContainer: tt.fields.groupContainer,
			}
			c.AddGroup(tt.in)
			got := c.groupContainer.groups
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoItemContainer() groups = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoItemContainer_GetTasks(t *testing.T) {
	type fields struct {
		taskContainer  *TaskContainer
		groupContainer *GroupContainer
	}
	tests := map[string]struct {
		fields fields
		want   []*Task
	}{
		"Success_ZeroTask": {
			fields{NewTaskContainer(), NewGroupContainer()},
			[]*Task{},
		},
		"Success_OneTask": {
			fields{
				&TaskContainer{
					tasks: []*Task{{"Task1", false}},
				},
				NewGroupContainer(),
			},
			[]*Task{{"Task1", false}},
		},
		"Success_ThreeTasks": {
			fields{
				&TaskContainer{
					tasks: []*Task{
						{"Task1", false},
						{"Task2", false},
						{"Task3", true},
					},
				},
				NewGroupContainer(),
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
				taskContainer:  tt.fields.taskContainer,
				groupContainer: tt.fields.groupContainer,
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
		taskContainer  *TaskContainer
		groupContainer *GroupContainer
	}
	tests := map[string]struct {
		fields fields
		want   []*Group
	}{
		"Success_ZeroGroup": {
			fields{
				NewTaskContainer(),
				NewGroupContainer(),
			},
			[]*Group{},
		},
		"Success_OneGroup": {
			fields{
				NewTaskContainer(),
				&GroupContainer{groups: []*Group{
					{"Group1", make([]*Task, 0)},
				}},
			},
			[]*Group{
				{"Group1", make([]*Task, 0)},
			},
		},
		"Success_ThreeGroups": {
			fields{
				NewTaskContainer(),
				&GroupContainer{groups: []*Group{
					{"Group1", make([]*Task, 0)},
					{"Group2", make([]*Task, 0)},
					{"Group3", make([]*Task, 0)},
				}},
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
				taskContainer:  tt.fields.taskContainer,
				groupContainer: tt.fields.groupContainer,
			}
			got := c.GetGroups()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
