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
			want: &TodoList{items: make([]Item, 0)},
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

func TestTodoList_AddItem(t *testing.T) {
	tests := map[string]struct {
		in   Item
		want []Item
	}{
		"Success_Task": {
			in:   &Task{"T", false},
			want: []Item{&Task{"T", false}},
		},
		"Success_Group": {
			in:   &Group{"G", make([]Item, 0)},
			want: []Item{&Group{"G", make([]Item, 0)}},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			list := TodoList{make([]Item, 0)}
			list.AddItem(tt.in)
			got := list.items
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddItem() items = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoList_GetItems(t *testing.T) {
	tests := map[string]struct {
		fields []Item
		want   []Item
	}{
		"ZeroItem": {
			fields: []Item{},
			want:   []Item{},
		},
		"TasksAndGroups": {
			fields: []Item{
				&Task{"Task1", false},
				&Task{"Task2", false},
				&Group{"Group1", []Item{&Task{"Task3", false}}},
				&Group{"Group2", []Item{&Task{"Task4", false}}},
			},
			want: []Item{
				&Task{"Task1", false},
				&Task{"Task2", false},
				&Group{"Group1", []Item{&Task{"Task3", false}}},
				&Group{"Group2", []Item{&Task{"Task4", false}}},
			},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			list := TodoList{items: tt.fields}
			got := list.GetItems()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItems() = %v, want %v", got, tt.want)
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
		"100%_Mix": {
			in: input{
				tasks:      []bool{true, true},
				groupTasks: [][]bool{{true, true}, {true, true}},
			},
			want: 1,
		},
		"50%_Mix_010101": {
			in: input{
				tasks:      []bool{false, true},
				groupTasks: [][]bool{{false, true}, {false, true}},
			},
			want: 0.5,
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
				c.AddItem(task)
			}
			for _, areTasksDone := range tt.in.groupTasks {
				group := &Group{}
				for _, isDone := range areTasksDone {
					group.AddItem(&Task{isDone: isDone})
				}
				c.AddItem(group)
			}

			got := c.Progress()

			if got != tt.want {
				t.Errorf("Progress() = %v, want %v", got, tt.want)
			}
		})
	}
}
