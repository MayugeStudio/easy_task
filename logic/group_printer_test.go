package logic

import (
	"bytes"
	"github.com/MayugeStudio/easy_task/domain"
	"testing"
)

func TestPrintGroups(t *testing.T) {
	tests := map[string]struct {
		in      []*domain.Group
		wantW   string
		wantErr bool
	}{
		"Success_1Group": {
			in: []*domain.Group{{"G", []*domain.Task{{"T1", false}, {"T2", true}}}},
			wantW: "" +
				"G\n" +
				"  [ ] T1\n" +
				"  [X] T2\n" +
				"  [##########          ]50.0%\n",
			wantErr: false,
		},
		"Success_2Group": {
			in: []*domain.Group{
				{"G1", []*domain.Task{{"T1", false}, {"T2", true}}},
				{"G2", []*domain.Task{{"T1", false}, {"T2", false}}},
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
			err := PrintGroups(w, tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrintGroups() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("PrintGroups() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_getGroupString(t *testing.T) {
	tests := map[string]struct {
		in   *domain.Group
		want string
	}{
		"100%": {
			in: &domain.Group{Title: "GroupTitle", Tasks: []*domain.Task{{"Task1", true}, {"Task2", true}}},
			want: "" +
				"GroupTitle\n" +
				"  [X] Task1\n" +
				"  [X] Task2\n" +
				"  [####################]100.0%\n",
		},
		"50%": {
			in: &domain.Group{Title: "GroupTitle", Tasks: []*domain.Task{{"Task1", true}, {"Task2", false}}},
			want: "" +
				"GroupTitle\n" +
				"  [X] Task1\n" +
				"  [ ] Task2\n" +
				"  [##########          ]50.0%\n",
		},
		"0%": {
			in: &domain.Group{Title: "GroupTitle", Tasks: []*domain.Task{{"Task1", false}, {"Task2", false}}},
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
