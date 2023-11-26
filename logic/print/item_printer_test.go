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
	type input struct {
		item domain.Item
	}
	tests := map[string]struct {
		in   input
		want string
	}{
		"Task_Done": {
			in:   input{item: newTask("TaskTitle", true)},
			want: "[X] TaskTitle\n",
		},
		"Task_Undone": {
			in:   input{item: newTask("TaskTitle", false)},
			want: "[ ] TaskTitle\n",
		},
		"Group": {
			in: input{item: newGroup("GroupTitle", []domain.Item{newTask("TaskTitle1", true), newTask("TaskTitle2", false)})},
			want: "" +
				"GroupTitle\n" +
				"  [X] TaskTitle1\n" +
				"  [ ] TaskTitle2\n" +
				"  [##########          ]50.0%\n",
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := getItemString(tt.in.item)
			if got != tt.want {
				t.Errorf("getItemString() gotW = %v, want %v", got, tt.want)
			}
		})
	}
}
