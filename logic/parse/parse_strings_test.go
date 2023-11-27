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

func TestToItems(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in   []string
		want domain.Items
	}{
		"Tasks": {
			in: []string{
				"- [ ] Task1",
				"- [X] Task2",
			},
			want: domain.Items{
				newTask("Task1", false),
				newTask("Task2", true),
			},
		},
		"Group": {
			in: []string{
				"- Group",
				"  - [ ] Task1",
				"  - [X] Task2",
			},
			want: []domain.Item{
				newGroup(
					"Group",
					[]domain.Item{
						newTask("Task1", false),
						newTask("Task2", true),
					},
				),
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
		//"NestedGroups": {
		//	in: []string{
		//		"- Group1",
		//		"  - Group1-1",
		//		"    - [X] Task1",
		//		"    - [ ] Task2",
		//		"  - Group1-2",
		//		"    - [X] Task1",
		//		"    - [ ] Task2",
		//		"- Group2",
		//		"  - Group2-1",
		//		"    - [X] Task1",
		//		"    - [ ] Task2",
		//		"  - Group2-2",
		//		"    - [X] Task1",
		//		"    - [ ] Task2",
		//	},
		//	want: []domain.Item{
		//		newGroup(
		//			"Group1",
		//			[]domain.Item{
		//				newGroup(
		//					"Group1-1",
		//					[]domain.Item{
		//						newTask("Task1", true),
		//						newTask("Task2", false),
		//					},
		//				),
		//				newGroup(
		//					"Group1-2",
		//					[]domain.Item{
		//						newTask("Task1", true),
		//						newTask("Task2", false),
		//					},
		//				),
		//			},
		//		),
		//		newGroup(
		//			"Group2",
		//			[]domain.Item{
		//				newGroup(
		//					"Group2-1",
		//					[]domain.Item{
		//						newTask("Task1", true),
		//						newTask("Task2", false),
		//					},
		//				),
		//				newGroup(
		//					"Group2-2",
		//					[]domain.Item{
		//						newTask("Task1", true),
		//						newTask("Task2", false),
		//					},
		//				),
		//			},
		//		),
		//	},
		//},
		"ContainInvalidString": {
			[]string{
				"- Group1",
				"  - [ ]Task1",
				"InvalidTaskString",
				"  - [X]Task2",
				"- Group2",
				"  - [ ]Task1",
				"    InvalidTaskString",
				"  - [X]Task2",
			},
			[]domain.Item{
				newGroup("Group1", []domain.Item{newTask("Task1", false), newTask("Task2", true)}),
				newGroup("Group2", []domain.Item{newTask("Task1", false), newTask("Task2", true)}),
			},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := ToItems(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				fmt.Println("Got")
				debug(got, 0)
				fmt.Println("Want")
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
		"Done": {
			in:   "- [X] TaskName",
			want: newTask("TaskName", true),
		},
		"Undone": {
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
		"GroupTitle":           {in: "- GroupName", want: newGroup("GroupName", make([]domain.Item, 0))},
		"GroupTitle_2Indented": {in: "  - GroupName", want: newGroup("GroupName", make([]domain.Item, 0))},
		"GroupTitle_4Indented": {in: "    - GroupName", want: newGroup("GroupName", make([]domain.Item, 0))},
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
