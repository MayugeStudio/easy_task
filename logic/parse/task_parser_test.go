package parse

import (
	"reflect"
	"testing"

	"github.com/MayugeStudio/easy_task/domain"
)

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
		"2Indent": {
			in:   "  - [ ] TaskName",
			want: newTask("TaskName", false),
		},
		"4Indent": {
			in:   "    - [ ] TaskName",
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
