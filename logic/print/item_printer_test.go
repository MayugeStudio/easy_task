package print

import (
	"bytes"
	"testing"

	"github.com/MayugeStudio/easy_task/domain"
)

func TestItems(t *testing.T) {
	tests := map[string]struct {
		in      []domain.Item
		wantW   string
		wantErr bool
	}{
		"Success_Tasks": {
			in: []domain.Item{newTask("T1", false), newTask("T2", true)},
			wantW: "" +
				"[ ] T1\n" +
				"[X] T2\n",
			wantErr: false,
		},
		"Success_Groups": {
			in: []domain.Item{
				newGroup("G1", []domain.Item{newTask("T1", false), newTask("T2", true)}),
				newGroup("G2", []domain.Item{newTask("T1", false), newTask("T2", false)}),
			},
			wantW: "" +
				"G1 [##########          ]50.0%\n" +
				"  [ ] T1\n" +
				"  [X] T2\n" +
				"G2 [                    ]0.0%\n" +
				"  [ ] T1\n" +
				"  [ ] T2\n",
			wantErr: false,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			w := &bytes.Buffer{}
			err := Items(w, tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Items() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Items() gotW = \n%v, want \n%v", gotW, tt.wantW)
			}
		})
	}
}

func Test_getItemString(t *testing.T) {
	tests := map[string]struct {
		in   domain.Item
		want string
	}{
		"Task_Done": {
			in:   newTask("Task", true),
			want: "[X] Task\n",
		},
		"Task_Undone": {
			in:   newTask("Task", false),
			want: "[ ] Task\n",
		},
		"Group": {
			in: newGroup("Group",
				[]domain.Item{
					newTask("Task1", true),
					newTask("Task2", false),
				},
			),
			want: "" +
				"Group [##########          ]50.0%\n" +
				"  [X] Task1\n" +
				"  [ ] Task2\n",
		},
		"NestedGroup": {
			in: newGroup("Group",
				[]domain.Item{
					newGroup("NestedGroup",
						[]domain.Item{
							newTask("Task1", true),
							newTask("Task2", false),
						},
					),
				},
			),
			want: "" +
				"Group [##########          ]50.0%\n" +
				"  NestedGroup [##########          ]50.0%\n" +
				"    [X] Task1\n" +
				"    [ ] Task2\n",
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := getItemString(tt.in, 0)
			if got != tt.want {
				t.Errorf("getItemString() gotW = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}
