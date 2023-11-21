package code

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
		"Success": {
			args{"TaskTitle", false},
			&Task{"TaskTitle", false}},
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

func TestNewTaskContainer(t *testing.T) {
	tests := map[string]struct {
		want *TaskContainer
	}{
		"Success": {&TaskContainer{tasks: make([]*Task, 0), doneTaskCount: 0}},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewTaskContainer()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskContainer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskContainer_AddTask(t *testing.T) {
	type fields struct {
		tasks         []*Task
		doneTaskCount int
	}
	tests := map[string]struct {
		fields fields
		in     *Task
		want   fields
	}{
		"AddDoneTask": {
			fields{make([]*Task, 0), 0},
			&Task{"TaskTitle", true},
			fields{[]*Task{{"TaskTitle", true}}, 1},
		},
		"AddUndoneTask": {
			fields{make([]*Task, 0), 0},
			&Task{"TaskTitle", false},
			fields{[]*Task{{"TaskTitle", false}}, 0},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			c := &TaskContainer{
				tasks:         tt.fields.tasks,
				doneTaskCount: tt.fields.doneTaskCount,
			}
			c.AddTask(tt.in)
			gotTasks := c.tasks
			gotDoneTaskCount := c.doneTaskCount

			if !reflect.DeepEqual(gotTasks, tt.want.tasks) {
				t.Errorf("TaskContainer() tasks = %v, want %v", gotTasks, tt.want.tasks)
			}
			if !reflect.DeepEqual(gotDoneTaskCount, tt.want.doneTaskCount) {
				t.Errorf("TaskContainer() doneTaskCount = %v, want %v", gotDoneTaskCount, tt.want.doneTaskCount)
			}
		})
	}
}
