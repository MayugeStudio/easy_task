package parse

import (
	"reflect"
	"testing"

	"github.com/MayugeStudio/easy_task/domain"
)

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
