package code

import (
	"reflect"
	"testing"
)

func TestParseStringsToTasks(t *testing.T) {
	tests := map[string]struct {
		in   []string
		want []*Task
	}{
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
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := ParseStringsToTasks(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseStringsToTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}
