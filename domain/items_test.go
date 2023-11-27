package domain

import (
	"reflect"
	"testing"
)

func TestNewItems(t *testing.T) {
	tests := map[string]struct {
		want Items
	}{
		"Success": {
			want: Items{},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewItems()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItems_Progress(t *testing.T) {
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
			items := Items{}
			for _, status := range tt.in.tasks {
				task := &Task{isDone: status}
				items = append(items, task)
			}
			for _, areTasksDone := range tt.in.groupTasks {
				group := &Group{}
				for _, isDone := range areTasksDone {
					group.AddItem(&Task{isDone: isDone})
				}
				items = append(items, group)
			}

			got := items.Progress()

			if got != tt.want {
				t.Errorf("Progress() = %v, want %v", got, tt.want)
			}
		})
	}
}
