package logic

import (
	"easy_task/src/domain"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func ConvertTaskPtrSliceToTaskValueSlice(S []*domain.Task) []domain.Task {
	result := make([]domain.Task, 0)
	for _, p := range S {
		result = append(result, *p)
	}
	return result
}

func ConvertGroupPtrSliceToGroupValueSlice(S []*domain.Group) string {
	var b strings.Builder
	for _, p := range S {
		b.WriteRune('[')
		b.WriteString("Title: " + p.Title)
		b.WriteString("  ")
		b.WriteString("tasks: " + fmt.Sprintf("%v", ConvertTaskPtrSliceToTaskValueSlice(p.Tasks)))
		b.WriteRune(']')
		b.WriteString("\n")
	}
	return b.String()
}

func TestParseStringsToTasks_OnlyTask(t *testing.T) {
	tests := map[string]struct {
		in   []string
		want []*domain.Task
	}{
		"DoneTasks": {
			[]string{
				"- [X]Task1",
				"- [X]Task2",
			},
			[]*domain.Task{
				{"Task1", true},
				{"Task2", true},
			},
		},
		"DoneTasks_Lowercase": {
			[]string{
				"- [x] Task1",
				"- [x] Task2",
			},
			[]*domain.Task{
				{"Task1", true},
				{"Task2", true},
			},
		},
		"UndoneTasks": {
			[]string{
				"- [ ] Task1",
				"- [ ] Task2",
			},
			[]*domain.Task{
				{"Task1", false},
				{"Task2", false},
			},
		},
		"MixPattern": {
			[]string{
				"- [ ] Task1",
				"- [X] Task2",
			},
			[]*domain.Task{
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
			[]*domain.Task{
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
		want []*domain.Group
	}{
		"DoneTasks": {
			[]string{
				"- TaskGroup",
				"  - [X]Task1",
				"  - [X]Task2",
			},
			[]*domain.Group{
				{
					"TaskGroup",
					[]*domain.Task{
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
			[]*domain.Group{
				{
					"TaskGroup",
					[]*domain.Task{
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
			[]*domain.Group{
				{
					"TaskGroup",
					[]*domain.Task{
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
			[]*domain.Group{
				{
					"TaskGroup",
					[]*domain.Task{
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
			[]*domain.Group{
				{
					"TaskGroup",
					[]*domain.Task{
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
			[]*domain.Group{
				{
					"TaskGroup",
					[]*domain.Task{
						{"Task1", false},
					},
				},
			},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := ParseStringsToTasks(tt.in)
			if !reflect.DeepEqual(got.GetGroups(), tt.want) {
				gotV := ConvertGroupPtrSliceToGroupValueSlice(got.GetGroups())
				wantV := ConvertGroupPtrSliceToGroupValueSlice(tt.want)
				t.Errorf("ParseStringsToTasks() = %v, want %v", gotV, wantV)
			}
		})
	}
}

func TestParseStringsToTasks_MultiGroupTask(t *testing.T) {
	tests := map[string]struct {
		in   []string
		want []*domain.Group
	}{
		"MixPattern_Multi": {
			[]string{
				"- TaskGroup1",
				"  - [ ]Task1",
				"  - [X]Task2",
				"- TaskGroup2",
				"  - [ ]Task1",
				"  - [X]Task2",
			},
			[]*domain.Group{
				{
					"TaskGroup1",
					[]*domain.Task{
						{"Task1", false},
						{"Task2", true},
					},
				},
				{
					"TaskGroup2",
					[]*domain.Task{
						{"Task1", false},
						{"Task2", true},
					},
				},
			},
		},
		"ContainInvalidTaskString": {
			[]string{
				"- TaskGroup1",
				"  - [ ]Task1",
				"  InvalidTaskString",
				"  - [X]Task2",
				"- TaskGroup2",
				"  - [ ]Task1",
				"  InvalidTaskString",
				"  - [X]Task2",
			},
			[]*domain.Group{
				{
					"TaskGroup1",
					[]*domain.Task{
						{"Task1", false},
						{"Task2", true},
					},
				},
				{
					"TaskGroup2",
					[]*domain.Task{
						{"Task1", false},
						{"Task2", true},
					},
				},
			},
		},
		"ContainInvalidTaskString_BadIndent": {
			[]string{
				"- TaskGroup1",
				"  - [ ]Task1",
				"InvalidTaskString",
				"  - [X]Task2",
				"- TaskGroup2",
				"  - [ ]Task1",
				"InvalidTaskString",
				"  - [X]Task2",
			},
			[]*domain.Group{
				{
					"TaskGroup1",
					[]*domain.Task{
						{"Task1", false},
					},
				},
				{
					"TaskGroup2",
					[]*domain.Task{
						{"Task1", false},
					},
				},
			},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := ParseStringsToTasks(tt.in)
			if !reflect.DeepEqual(got.GetGroups(), tt.want) {
				gotV := ConvertGroupPtrSliceToGroupValueSlice(got.GetGroups())
				wantV := ConvertGroupPtrSliceToGroupValueSlice(tt.want)
				t.Errorf("ParseStringsToTasks() = %v, want %v", gotV, wantV)
			}
		})
	}
}

func Test_parseSingleTaskString(t *testing.T) {
	tests := map[string]struct {
		in   string
		want *domain.Task
	}{"ValidSingleTaskString_Done": {
		"- [X] TaskName",
		&domain.Task{
			Title:  "TaskName",
			IsDone: true,
		},
	},
		"ValidSingleTaskString_Undone": {
			"- [ ] TaskName",
			&domain.Task{
				Title:  "TaskName",
				IsDone: false,
			},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := parseTaskString(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTaskString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseGroupTaskTitle(t *testing.T) {
	tests := map[string]struct {
		in   string
		want *domain.Group
	}{
		"ValidGroupTitle": {
			"- GroupName",
			&domain.Group{
				Title: "GroupName",
				Tasks: make([]*domain.Task, 0),
			},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := parseGroupTaskTitle(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseGroupTaskTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}