package logic

import (
	"bytes"
	"testing"

	"github.com/MayugeStudio/easy_task/domain"
)

func TestPrintTasks(t *testing.T) {
	tests := map[string]struct {
		in      []*domain.Task
		wantW   string
		wantErr bool
	}{
		"Success_1Task": {
			in:      []*domain.Task{{"Task1", false}},
			wantW:   "[ ] Task1\n",
			wantErr: false,
		},
		"Success_3Tasks": {
			in: []*domain.Task{{"Task1", false}, {"Task2", true}, {"Task3", true}},
			wantW: "" +
				"[ ] Task1\n" +
				"[X] Task2\n" +
				"[X] Task3\n",
			wantErr: false,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			w := &bytes.Buffer{}
			err := PrintTasks(w, tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrintTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotW := w.String()
			if gotW != tt.wantW {
				t.Errorf("PrintTasks() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestPrintGroups(t *testing.T) {
	tests := map[string]struct {
		in      []*domain.Group
		wantW   string
		wantErr bool
	}{
		"Success_1Group": {
			in: []*domain.Group{{"GroupTitle", []*domain.Task{{"Task1", false}, {"Task2", true}}}},
			wantW: "" +
				"GroupTitle\n" +
				"  [ ] Task1\n" +
				"  [X] Task2\n" +
				"  [##########          ]50%\n",
			wantErr: false,
		},
		"Success_3Group": {
			in: []*domain.Group{
				{"GroupTitle1", []*domain.Task{{"Task1", false}, {"Task2", true}}},
				{"GroupTitle2", []*domain.Task{{"Task1", false}, {"Task2", false}}},
				{"GroupTitle3", []*domain.Task{{"Task1", true}, {"Task2", true}}},
			},
			wantW: "" +
				"GroupTitle1\n" +
				"  [ ] Task1\n" +
				"  [X] Task2\n" +
				"  [##########          ]50%\n" +
				"GroupTitle2\n" +
				"  [ ] Task1\n" +
				"  [ ] Task2\n" +
				"  [                    ]0%\n" +
				"GroupTitle3\n" +
				"  [X] Task1\n" +
				"  [X] Task2\n" +
				"  [####################]100%\n",
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

func TestPrintProgress(t *testing.T) {
	tests := map[string]struct {
		in      []*domain.Task
		wantW   string
		wantErr bool
	}{
		"100%": {
			in:      []*domain.Task{{"Task1", true}, {"Task2", true}},
			wantW:   "[########################################]100%",
			wantErr: false,
		},
		"50%": {
			in:      []*domain.Task{{"Task1", true}, {"Task2", false}},
			wantW:   "[####################                    ]50%",
			wantErr: false,
		},
		"0%": {
			in:      []*domain.Task{{"Task1", false}, {"Task2", false}},
			wantW:   "[                                        ]0%",
			wantErr: false,
		},
		"NonTask": {
			in:      []*domain.Task{},
			wantW:   "[                                        ]0%",
			wantErr: false,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			w := &bytes.Buffer{}
			c := domain.NewTodoList()
			for _, task := range tt.in {
				c.AddTask(task)
			}
			err := PrintProgress(w, c)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrintProgress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("PrintProgress() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_getTaskString(t *testing.T) {
	type input struct {
		task      *domain.Task
		maxLength int
	}
	tests := map[string]struct {
		in   input
		want string
	}{
		"Success_Done":   {in: input{task: &domain.Task{Title: "TaskTitle", IsDone: true}, maxLength: 10}, want: "[X] TaskTitle "},
		"Success_Undone": {in: input{task: &domain.Task{Title: "TaskTitle", IsDone: false}, maxLength: 10}, want: "[ ] TaskTitle "},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := getTaskString(tt.in.task, tt.in.maxLength)
			if got != tt.want {
				t.Errorf("getTaskString() gotW = %v, want %v", got, tt.want)
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
				"  [####################]100%\n",
		},
		"50%": {
			in: &domain.Group{Title: "GroupTitle", Tasks: []*domain.Task{{"Task1", true}, {"Task2", false}}},
			want: "" +
				"GroupTitle\n" +
				"  [X] Task1\n" +
				"  [ ] Task2\n" +
				"  [##########          ]50%\n",
		},

		"0%": {
			in: &domain.Group{Title: "GroupTitle", Tasks: []*domain.Task{{"Task1", false}, {"Task2", false}}},
			want: "" +
				"GroupTitle\n" +
				"  [ ] Task1\n" +
				"  [ ] Task2\n" +
				"  [                    ]0%\n",
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

func Test_getProgressString(t *testing.T) {
	type input struct {
		progress float64
		length   float64
	}
	tests := map[string]struct {
		in   input
		want string
	}{
		"100%": {
			in:   input{progress: 1, length: 40},
			want: "[########################################]100%",
		},
		"50%": {
			in:   input{progress: 0.5, length: 40},
			want: "[####################                    ]50%",
		},
		"25%": {
			in:   input{progress: 0.25, length: 40},
			want: "[##########                              ]25%",
		},
		"0%": {
			in:   input{progress: 0, length: 40},
			want: "[                                        ]0%",
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := getProgressString(tt.in.progress, tt.in.length)
			if got != tt.want {
				t.Errorf("getProgressString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMaxTaskNameLength(t *testing.T) {
	tests := map[string]struct {
		in   []*domain.Task
		want int
	}{
		"Success_Length5": {
			in: []*domain.Task{
				{"12", false},
				{"123", false},
				{"12345", false},
			},
			want: 5,
		},
		"Success_Length10": {
			in: []*domain.Task{
				{"1234567890", false},
				{"1234567", false},
				{"123", false},
			},
			want: 10,
		},
		"Success_Length20": {
			in: []*domain.Task{
				{"12345678901234567890", false},
				{"123456789012", false},
				{"123", false},
				{"1234567", false},
			},
			want: 20,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := getMaxTaskNameLength(tt.in); got != tt.want {
				t.Errorf("getMaxTaskNameLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
