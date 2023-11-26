package parse

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/MayugeStudio/easy_task/domain"
)

var newTask = domain.NewTask
var newGroup = func(title string, tasks []*domain.Task) *domain.Group {
	g := domain.NewGroup(title)
	for _, task := range tasks {
		g.AddTask(task)
	}
	return g
}

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
		b.WriteString("Title: " + p.Title())
		b.WriteString("  ")
		b.WriteString("tasks: " + fmt.Sprintf("%v", ConvertTaskPtrSliceToTaskValueSlice(p.Tasks())))
		b.WriteRune(']')
		b.WriteString("\n")
	}
	return b.String()
}

func TestStringsToTasks_OnlyTask(t *testing.T) {
	tests := map[string]struct {
		in   []string
		want []*domain.Task
	}{
		"DoneTasks": {
			in:   []string{"- [X]Task1", "- [X]Task2"},
			want: []*domain.Task{newTask("Task1", true), newTask("Task2", true)},
		},
		"DoneTasks_Lowercase": {
			in:   []string{"- [x] Task1", "- [x] Task2"},
			want: []*domain.Task{newTask("Task1", true), newTask("Task2", true)},
		},
		"UndoneTasks": {
			in:   []string{"- [ ] Task1", "- [ ] Task2"},
			want: []*domain.Task{newTask("Task1", false), newTask("Task2", false)},
		},
		"MixPattern": {
			in:   []string{"- [ ] Task1", "- [X] Task2"},
			want: []*domain.Task{newTask("Task1", false), newTask("Task2", true)},
		},
		"ContainInvalidTaskString": {
			in:   []string{"- [ ] Task1", "InvalidTaskString", "- [X] Task2"},
			want: []*domain.Task{newTask("Task1", false), newTask("Task2", true)},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := ToTodoList(tt.in)
			if !reflect.DeepEqual(got.GetTasks(), tt.want) {
				gotV := ConvertTaskPtrSliceToTaskValueSlice(got.GetTasks())
				wantV := ConvertTaskPtrSliceToTaskValueSlice(tt.want)
				t.Errorf("ToTodoList() = %v, want %v", gotV, wantV)
			}
		})
	}
}

func TestStringsToTasks_OnlyGroupTask(t *testing.T) { // FIXME: Rename function name 'OnlyGroupTask' to 'OnlySingleGroup'.
	tests := map[string]struct {
		in   []string
		want []*domain.Group
	}{
		"DoneTasks": {
			in: []string{"- TaskGroup", "  - [X]Task1", "  - [X]Task2"},
			want: []*domain.Group{newGroup(
				"TaskGroup",
				[]*domain.Task{newTask("Task1", true), newTask("Task2", true)}),
			},
		},
		"DoneTasks_Lowercase": {
			in: []string{"- TaskGroup", "  - [x]Task1", "  - [x]Task2"},
			want: []*domain.Group{newGroup(
				"TaskGroup",
				[]*domain.Task{newTask("Task1", true), newTask("Task2", true)}),
			},
		},
		"UndoneTasks": {
			in: []string{"- TaskGroup", "  - [ ]Task1", "  - [ ]Task2"},
			want: []*domain.Group{newGroup(
				"TaskGroup",
				[]*domain.Task{newTask("Task1", false), newTask("Task2", false)}),
			},
		},
		"MixPattern": {
			in: []string{"- TaskGroup", "  - [ ]Task1", "  - [X]Task2"},
			want: []*domain.Group{newGroup(
				"TaskGroup",
				[]*domain.Task{newTask("Task1", false), newTask("Task2", true)}),
			},
		},
		"ContainInvalidTaskString": {
			in: []string{"- TaskGroup", "  - [ ]Task1", "  InvalidTaskString", "  - [X]Task2"},
			want: []*domain.Group{newGroup(
				"TaskGroup",
				[]*domain.Task{newTask("Task1", false), newTask("Task2", true)}),
			},
		},
		"ContainInvalidTaskString_BadIndent": {
			in: []string{"- TaskGroup", "  - [ ]Task1", "InvalidTaskString", "  - [X]Task2"},
			want: []*domain.Group{newGroup(
				"TaskGroup",
				[]*domain.Task{newTask("Task1", false)}),
			},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := ToTodoList(tt.in)
			if !reflect.DeepEqual(got.GetGroups(), tt.want) {
				gotV := ConvertGroupPtrSliceToGroupValueSlice(got.GetGroups())
				wantV := ConvertGroupPtrSliceToGroupValueSlice(tt.want)
				t.Errorf("ToTodoList() = %v, want %v", gotV, wantV)
			}
		})
	}
}

func TestStringsToTasks_MultiGroupTask(t *testing.T) { // FIXME: Rename test function name.
	tests := map[string]struct {
		in   []string
		want []*domain.Group
	}{
		"MixPattern": {
			in: []string{
				"- TaskGroup1",
				"  - [ ]Task1",
				"  - [X]Task2",
				"- TaskGroup2",
				"  - [ ]Task1",
				"  - [X]Task2",
			},
			want: []*domain.Group{
				newGroup("TaskGroup1", []*domain.Task{newTask("Task1", false), newTask("Task2", true)}),
				newGroup("TaskGroup2", []*domain.Task{newTask("Task1", false), newTask("Task2", true)}),
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
				newGroup("TaskGroup1", []*domain.Task{newTask("Task1", false), newTask("Task2", true)}),
				newGroup("TaskGroup2", []*domain.Task{newTask("Task1", false), newTask("Task2", true)}),
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
				newGroup("TaskGroup1", []*domain.Task{newTask("Task1", false)}),
				newGroup("TaskGroup2", []*domain.Task{newTask("Task1", false)}),
			},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := ToTodoList(tt.in)
			if !reflect.DeepEqual(got.GetGroups(), tt.want) {
				gotV := ConvertGroupPtrSliceToGroupValueSlice(got.GetGroups())
				wantV := ConvertGroupPtrSliceToGroupValueSlice(tt.want)
				t.Errorf("ToTodoList() = %v, want %v", gotV, wantV)
			}
		})
	}
}

func Test_toTask(t *testing.T) {
	tests := map[string]struct {
		in   string
		want *domain.Task
	}{
		"ValidSingleTaskString_Done": {
			in:   "- [X] TaskName",
			want: newTask("TaskName", true),
		},
		"ValidSingleTaskString_Undone": {
			in:   "- [ ] TaskName",
			want: newTask("TaskName", false),
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := toTask(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toGroup(t *testing.T) {
	tests := map[string]struct {
		in   string
		want *domain.Group
	}{
		"ValidGroupTitle": {in: "- GroupName", want: newGroup("GroupName", make([]*domain.Task, 0))},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := toGroup(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
