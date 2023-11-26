package print

import (
	"bytes"
	"testing"

	"github.com/MayugeStudio/easy_task/domain"
)

func TestGroups(t *testing.T) { // FIXME: This test function's concern is not about group's and task's title.
	tests := map[string]struct {
		in      []*domain.Group
		wantW   string
		wantErr bool
	}{
		"Success_1Group": {
			in: []*domain.Group{newGroup("G", []*domain.Task{newTask("T1", false), newTask("T2", true)})},
			wantW: "" +
				"G\n" +
				"  [ ] T1\n" +
				"  [X] T2\n" +
				"  [##########          ]50.0%\n",
			wantErr: false,
		},
		"Success_2Group": {
			in: []*domain.Group{
				newGroup("G1", []*domain.Task{newTask("T1", false), newTask("T2", true)}),
				newGroup("G2", []*domain.Task{newTask("T1", false), newTask("T2", false)}),
			},
			wantW: "" +
				"G1\n" +
				"  [ ] T1\n" +
				"  [X] T2\n" +
				"  [##########          ]50.0%\n" +
				"G2\n" +
				"  [ ] T1\n" +
				"  [ ] T2\n" +
				"  [                    ]0.0%\n",
			wantErr: false,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			w := &bytes.Buffer{}
			err := Groups(w, tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Groups() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Groups() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_getGroupString(t *testing.T) { // FIXME: This test code's concern is not about group progress.
	tests := map[string]struct {
		in   *domain.Group
		want string
	}{
		"100%": {
			in: newGroup("GroupTitle", []*domain.Task{newTask("Task1", true), newTask("Task2", true)}),
			want: "" +
				"GroupTitle\n" +
				"  [X] Task1\n" +
				"  [X] Task2\n" +
				"  [####################]100.0%\n",
		},
		"50%": {
			in: newGroup("GroupTitle", []*domain.Task{newTask("Task1", true), newTask("Task2", false)}),
			want: "" +
				"GroupTitle\n" +
				"  [X] Task1\n" +
				"  [ ] Task2\n" +
				"  [##########          ]50.0%\n",
		},
		"0%": {
			in: newGroup("GroupTitle", []*domain.Task{newTask("Task1", false), newTask("Task2", false)}),
			want: "" +
				"GroupTitle\n" +
				"  [ ] Task1\n" +
				"  [ ] Task2\n" +
				"  [                    ]0.0%\n",
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := getGroupString(tt.in)
			if got != tt.want {
				t.Errorf("getGroupString() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}
