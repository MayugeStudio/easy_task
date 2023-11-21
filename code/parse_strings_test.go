package code

import (
	"reflect"
	"testing"
)

func ConvertTaskPtrSliceToTaskValueSlice(S []*Task) []Task {
	result := make([]Task, 0)
	for _, p := range S {
		result = append(result, *p)
	}
	return result
}

func ConvertGroupPtrSliceToGroupValueSlice(S []*Group) []Group {
	result := make([]Group, 0)
	for _, p := range S {
		result = append(result, *p)
	}
	return result
}

func TestParseStringsToTasks_OnlyTask(t *testing.T) {
	tests := map[string]struct {
		in   []string
		want []*Task
	}{
		"DoneTasks": {
			[]string{
				"- [X]Task1",
				"- [X]Task2",
			},
			[]*Task{
				{"Task1", true},
				{"Task2", true},
			},
		},
		"DoneTasks_Lowercase": {
			[]string{
				"- [x] Task1",
				"- [x] Task2",
			},
			[]*Task{
				{"Task1", true},
				{"Task2", true},
			},
		},
		"UndoneTasks": {
			[]string{
				"- [ ] Task1",
				"- [ ] Task2",
			},
			[]*Task{
				{"Task1", false},
				{"Task2", false},
			},
		},
		"MixPattern": {
			[]string{
				"- [ ] Task1",
				"- [X] Task2",
			},
			[]*Task{
				{"Task1", false},
				{"Task2", true},
			},
		},
		"ContainInvalidTaskString": {
			[]string{
				"- [ ] Task1",
				"InvalidTaskString",
				"- [X] Task2",
			},
			[]*Task{
				{"Task1", false},
				{"Task2", true},
			},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := ParseStringsToTasks(tt.in)
			if !reflect.DeepEqual(got.GetTasks(), tt.want) {
				gotV := ConvertTaskPtrSliceToTaskValueSlice(got.GetTasks())
				wantV := ConvertTaskPtrSliceToTaskValueSlice(tt.want)
				t.Errorf("ParseStringsToTasks() = %v, want %v", gotV, wantV)
			}
		})
	}
}

func TestParseStringsToTasks_OnlyGroupTask(t *testing.T) {
	tests := map[string]struct {
		in   []string
		want []*Group
	}{
		"DoneTasks": {
			[]string{
				"- TaskGroup",
				"  - [X]Task1",
				"  - [X]Task2",
			},
			[]*Group{
				{
					[]*Task{
						{"Task1", true},
						{"Task2", true},
					},
				},
			},
		},
		"DoneTasks_Lowercase": {
			[]string{
				"- TaskGroup",
				"  - [x]Task1",
				"  - [x]Task2",
			},
			[]*Group{
				{
					[]*Task{
						{"Task1", true},
						{"Task2", true},
					},
				},
			},
		},
		"UndoneTasks": {
			[]string{
				"- TaskGroup",
				"  - [ ]Task1",
				"  - [ ]Task2",
			},
			[]*Group{
				{
					[]*Task{
						{"Task1", false},
						{"Task2", false},
					},
				},
			},
		},
		"MixPattern": {
			[]string{
				"- TaskGroup",
				"  - [ ]Task1",
				"  - [X]Task2",
			},
			[]*Group{
				{
					[]*Task{
						{"Task1", false},
						{"Task2", true},
					},
				},
			},
		},
		"ContainInvalidTaskString": {
			[]string{
				"- TaskGroup",
				"  - [ ]Task1",
				"  InvalidTaskString",
				"  - [X]Task2",
			},
			[]*Group{
				{
					[]*Task{
						{"Task1", false},
						{"Task2", true},
					},
				},
			},
		},
		"ContainInvalidTaskString_BadIndent": {
			[]string{
				"- TaskGroup",
				"  - [ ]Task1",
				"InvalidTaskString",
				"  - [X]Task2",
			},
			[]*Group{
				{
					[]*Task{
						{"Task1", false},
					},
				},
			},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := ParseStringsToTasks(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				gotV := ConvertGroupPtrSliceToGroupValueSlice(got.GetGroups())
				wantV := ConvertGroupPtrSliceToGroupValueSlice(tt.want)
				t.Errorf("ParseStringsToTasks() = %v, want %v", gotV, wantV)
			}
		})
	}
}