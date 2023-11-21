package code

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func ConvertTaskPtrSliceToTaskValueSlice(S []*Task) []Task {
	result := make([]Task, 0)
	for _, p := range S {
		result = append(result, *p)
	}
	return result
}

func ConvertGroupPtrSliceToGroupValueSlice(S []*Group) string {
	var b strings.Builder
	for _, p := range S {
		b.WriteRune('[')
		b.WriteString("Title: " + p.title)
		b.WriteString("  ")
		b.WriteString("tasks: " + fmt.Sprintf("%v", ConvertTaskPtrSliceToTaskValueSlice(p.tasks)))
		b.WriteRune(']')
		b.WriteString("\n")
	}
	return b.String()
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
					"TaskGroup",
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
					"TaskGroup",
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
					"TaskGroup",
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
					"TaskGroup",
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
					"TaskGroup",
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
					"TaskGroup",
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
		want *Task
	}{"ValidSingleTaskString_Done": {
		"- [X] TaskName",
		&Task{
			Title:  "TaskName",
			IsDone: true,
		},
	},
		"ValidSingleTaskString_Undone": {
			"- [ ] TaskName",
			&Task{
				Title:  "TaskName",
				IsDone: false,
			},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := parseSingleTaskString(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSingleTaskString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseGroupTaskTitle(t *testing.T) {
	tests := map[string]struct {
		in   string
		want *Group
	}{
		"ValidGroupTitle": {
			"- GroupName",
			&Group{
				title: "GroupName",
				tasks: make([]*Task, 0),
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

func Test_parseGroupTaskString(t *testing.T) {
	tests := map[string]struct {
		in   string
		want *Task
	}{
		"ValidGroupTaskString_Done": {
			"  - [X] GroupTaskString",
			&Task{
				Title:  "GroupTaskString",
				IsDone: true,
			},
		},
		"ValidGroupTaskString_Undone": {
			"  - [ ] GroupTaskString",
			&Task{
				Title:  "GroupTaskString",
				IsDone: false,
			},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := parseGroupTaskString(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseGroupTaskString() = %v, want %v", got, tt.want)
			}
		})
	}
}
