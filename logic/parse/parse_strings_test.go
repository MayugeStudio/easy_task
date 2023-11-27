package parse

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/MayugeStudio/easy_task/domain"
)

var newTask = domain.NewTask
var newGroup = func(title string, items []domain.Item) *domain.Group {
	g := domain.NewGroup(title)
	for _, item := range items {
		g.AddItem(item)
	}
	return g
}

func debug(items []domain.Item, indent int) {
	for _, item := range items {
		fmt.Println(strings.Repeat(" ", indent), item.Title())
		if item.IsParent() {
			debug(item.Children(), indent+2)
		}
	}
}

func TestToItems_OnlyTask(t *testing.T) { // TODO: Refactor test cases.
	t.Parallel()
	tests := map[string]struct {
		in   []string
		want domain.Items
	}{
		"DoneTasks_Lowercase": {
			in:   []string{"- [x] Task1", "- [x] Task2"},
			want: domain.Items{newTask("Task1", true), newTask("Task2", true)},
		},
		"1Done1Undone": {
			in:   []string{"- [ ] Task1", "- [X] Task2"},
			want: domain.Items{newTask("Task1", false), newTask("Task2", true)},
		},
		"ContainInvalidTaskString": {
			in:   []string{"- [ ] Task1", "InvalidTaskString", "- [X] Task2"},
			want: domain.Items{newTask("Task1", false), newTask("Task2", true)},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := ToItems(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToItems_OnlySingleGroup(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in   []string
		want domain.Items
	}{
		"1GroupIn1DoneAnd1Undone": {
			in: []string{
				"- Group",
				"  - [ ] Task1",
				"  - [X] Task2",
			},
			want: []domain.Item{
				newGroup(
					"Group",
					[]domain.Item{newTask("Task1", false), newTask("Task2", true)},
				),
			},
		},
		"ContainInvalidTaskString": {
			in: []string{
				"- Group",
				"  - [ ] Task1",
				"  InvalidTaskString",
				"  - [X] Task2",
			},
			want: []domain.Item{
				newGroup(
					"Group",
					[]domain.Item{newTask("Task1", false), newTask("Task2", true)},
				),
			},
		},
		"ContainInvalidTaskString_BadIndent": {
			in: []string{
				"- Group",
				"  - [ ] Task1",
				"InvalidTaskString",
				"  - [X] Task2",
			},
			want: []domain.Item{
				newGroup(
					"Group",
					[]domain.Item{newTask("Task1", false)},
				),
			},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := ToItems(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToTodoList_MultiGroup(t *testing.T) { // FIXME: Rename test function name.
	t.Parallel()
	tests := map[string]struct {
		in   []string
		want domain.Items
	}{
		"Groups": {
			in: []string{
				"- Group1",
				"  - [ ]Task1",
				"  - [X]Task2",
				"- Group2",
				"  - [ ]Task1",
				"  - [X]Task2",
			},
			want: []domain.Item{
				newGroup("Group1", []domain.Item{newTask("Task1", false), newTask("Task2", true)}),
				newGroup("Group2", []domain.Item{newTask("Task1", false), newTask("Task2", true)}),
			},
		},
		"NestedGroup": {
			in: []string{
				"- Group1",
				"  - Group2",
				"    - [X] Task1",
				"    - [ ] Task2",
			},
			want: []domain.Item{
				newGroup(
					"Group1",
					[]domain.Item{
						newGroup(
							"Group2",
							[]domain.Item{newTask("Task1", true), newTask("Task2", false)}),
					}),
			},
		},
		"ContainInvalidTaskString": {
			[]string{
				"- Group1",
				"  - [ ]Task1",
				"  InvalidTaskString",
				"  - [X]Task2",
				"- Group2",
				"  - [ ]Task1",
				"  InvalidTaskString",
				"  - [X]Task2",
			},
			[]domain.Item{
				newGroup("Group1", []domain.Item{newTask("Task1", false), newTask("Task2", true)}),
				newGroup("Group2", []domain.Item{newTask("Task1", false), newTask("Task2", true)}),
			},
		},
		"ContainInvalidTaskString_BadIndent": {
			[]string{
				"- Group1",
				"  - [ ]Task1",
				"InvalidTaskString",
				"  - [X]Task2",
				"- Group2",
				"  - [ ]Task1",
				"InvalidTaskString",
				"  - [X]Task2",
			},
			[]domain.Item{
				newGroup("Group1", []domain.Item{newTask("Task1", false)}),
				newGroup("Group2", []domain.Item{newTask("Task1", false)}),
			},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := ToItems(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				debug(got, 0)
				debug(tt.want, 0)
				t.Errorf("ToItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toTask(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
	tests := map[string]struct {
		in   string
		want *domain.Group
	}{
		"ValidGroupTitle": {in: "- GroupName", want: newGroup("GroupName", make([]domain.Item, 0))},
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
