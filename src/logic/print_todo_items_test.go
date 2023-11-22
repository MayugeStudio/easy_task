package logic

import (
	"bytes"
	"easy_task/src/domain"
	"testing"
)

func TestPrintTasks(t *testing.T) {
	tests := map[string]struct {
		in      []*domain.Task
		wantW   string
		wantErr bool
	}{
		"Success_1Task": {
			[]*domain.Task{
				{"Task1", false},
			},
			"[ ] Task1\n",
			false,
		},
		"Success_3Tasks": {
			[]*domain.Task{
				{"Task1", false},
				{"Task2", true},
				{"Task3", true},
			},
			"[ ] Task1\n[X] Task2\n[X] Task3\n",
			false,
		},
		"Success_10Tasks": {
			[]*domain.Task{
				{"0Hi", false},
				{"1BuyTheMilk", false},
				{"2MaxLengthName", false},
				{"3ILikeSinging", true},
				{"4I'm Gopher", true},
			},
			"" +
				"[ ] 0Hi           \n" +
				"[ ] 1BuyTheMilk   \n" +
				"[ ] 2MaxLengthName\n" +
				"[X] 3ILikeSinging \n" +
				"[X] 4I'm Gopher   \n",
			false,
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
			if gotW := w.String(); gotW != tt.wantW {
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
			[]*domain.Group{
				{"GroupTitle",
					[]*domain.Task{
						{"Task1", false},
						{"Task2", true},
					},
				},
			},
			"" +
				"GroupTitle\n" +
				"  [ ] Task1\n" +
				"  [X] Task2\n" +
				"  [##########          ]50%\n",
			false,
		},
		"Success_3Group": {
			[]*domain.Group{
				{"GroupTitle1",
					[]*domain.Task{
						{"Task1", false},
						{"Task2", true},
					},
				},
				{"GroupTitle2",
					[]*domain.Task{
						{"Task1", false},
						{"Task2", false},
					},
				},
				{"GroupTitle3",
					[]*domain.Task{
						{"Task1", true},
						{"Task2", true},
					},
				},
			},
			"" +
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
			false,
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
			in: []*domain.Task{
				{"Task1", true},
				{"Task2", true},
			},
			wantW:   "[########################################]100%",
			wantErr: false,
		},
		"50%_01": {
			in: []*domain.Task{
				{"Task1", true},
				{"Task2", false},
			},
			wantW:   "[####################                    ]50%",
			wantErr: false,
		},
		"25%": {
			in: []*domain.Task{
				{"Task1", true},
				{"Task2", false},
				{"Task3", false},
				{"Task4", false},
			},
			wantW:   "[##########                              ]25%",
			wantErr: false,
		},
		"0%": {
			in: []*domain.Task{
				{"Task1", false},
				{"Task2", false},
			},
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
			c := domain.NewTodoItemContainer()
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
		"Success_Done": {
			input{
				&domain.Task{
					"TaskTitle", true,
				},
				10,
			},
			"[X] TaskTitle ",
		},
		"Success_Undone": {
			input{
				&domain.Task{
					"TaskTitle", false,
				},
				10,
			},
			"[ ] TaskTitle ",
		},
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
			in: &domain.Group{
				Title: "GroupTitle",
				Tasks: []*domain.Task{
					{"Task1", true},
					{"Task2", true},
				},
			},
			want: "" +
				"GroupTitle\n" +
				"  [X] Task1\n" +
				"  [X] Task2\n" +
				"  [####################]100%\n",
		},
		"50%": {
			in: &domain.Group{
				Title: "GroupTitle",
				Tasks: []*domain.Task{
					{"Task1", true},
					{"Task2", false},
				},
			},
			want: "" +
				"GroupTitle\n" +
				"  [X] Task1\n" +
				"  [ ] Task2\n" +
				"  [##########          ]50%\n",
		},

		"0%": {
			in: &domain.Group{
				Title: "GroupTitle",
				Tasks: []*domain.Task{
					{"Task1", false},
					{"Task2", false},
				},
			},
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
			input{
				1,
				40,
			},
			"[########################################]100%",
		},
		"50%": {
			input{
				0.5,
				40,
			},
			"[####################                    ]50%",
		},
		"25%": {
			input{
				0.25,
				40,
			},
			"[##########                              ]25%",
		},
		"0%": {
			input{
				0,
				40,
			},
			"[                                        ]0%",
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
			[]*domain.Task{
				{"12", false},
				{"123", false},
				{"12345", false},
			},
			5,
		},
		"Success_Length10": {
			[]*domain.Task{
				{"1234567890", false},
				{"1234567", false},
				{"123", false},
			},
			10,
		},
		"Success_Length20": {
			[]*domain.Task{
				{"12345678901234567890", false},
				{"123456789012", false},
				{"123", false},
				{"1234567", false},
			},
			20,
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

func Test_calculateProgress(t *testing.T) {
	type input struct {
		tasks  []*domain.Task
		groups []*domain.Group
	}
	tests := map[string]struct {
		in   input
		want float64
	}{
		"100%": {
			in: input{
				tasks: []*domain.Task{
					{"Task1", true},
					{"Task2", true},
				},
				groups: []*domain.Group{
					{
						"Group1",
						[]*domain.Task{
							{"Task1", true},
						},
					},
					{
						"Group2",
						[]*domain.Task{
							{"Task1", true},
							{"Task2", true},
						},
					},
				},
			},
			want: 1.0,
		},
		"100%_NoTask": {
			in: input{
				tasks: []*domain.Task{},
				groups: []*domain.Group{
					{
						"Group1",
						[]*domain.Task{
							{"Task1", true},
						},
					},
					{
						"Group2",
						[]*domain.Task{
							{"Task1", true},
							{"Task2", true},
						},
					},
				},
			},
			want: 1.0,
		},
		"100%_NoGroup": {
			in: input{
				tasks: []*domain.Task{
					{"Task1", true},
					{"Task2", true},
				},
				groups: []*domain.Group{},
			},
			want: 1.0,
		},
		"50%": {
			in: input{
				tasks: []*domain.Task{
					{"Task1", true},
				},
				groups: []*domain.Group{
					{
						"Group1",
						[]*domain.Task{
							{"Task1", true},
						},
					},
					{
						"Group2",
						[]*domain.Task{
							{"Task1", false},
							{"Task2", false},
						},
					},
				},
			},
			want: 0.5,
		},
		"50%_NoTask": {
			in: input{
				tasks: []*domain.Task{
					{"Task1", true},
				},
				groups: []*domain.Group{
					{
						"Group1",
						[]*domain.Task{
							{"Task1", true},
						},
					},
					{
						"Group2",
						[]*domain.Task{
							{"Task1", false},
							{"Task2", false},
						},
					},
				},
			},
			want: 0.5,
		},
		"50%_NoGroup": {
			in: input{
				tasks: []*domain.Task{
					{"Task1", true},
				},
				groups: []*domain.Group{
					{
						"Group1",
						[]*domain.Task{
							{"Task1", true},
						},
					},
					{
						"Group2",
						[]*domain.Task{
							{"Task1", false},
							{"Task2", false},
						},
					},
				},
			},
			want: 0.5,
		},
		"0%": {
			in: input{
				tasks: []*domain.Task{
					{"Task1", false},
					{"Task2", false},
				},
				groups: []*domain.Group{
					{
						"Group1",
						[]*domain.Task{
							{"Task1", false},
							{"Task2", false},
							{"Task3", false},
						},
					},
					{
						"Group2",
						[]*domain.Task{
							{"Task1", false},
							{"Task2", false},
						},
					},
				},
			},
			want: 0.0,
		},
		"0%_NoTask": {
			in: input{
				tasks: []*domain.Task{},
				groups: []*domain.Group{
					{
						"Group1",
						[]*domain.Task{
							{"Task1", false},
						},
					},
					{
						"Group2",
						[]*domain.Task{
							{"Task1", false},
							{"Task2", false},
						},
					},
				},
			},
			want: 0.0,
		},
		"0%_NoGroup": {
			in: input{
				tasks: []*domain.Task{
					{"Task1", false},
					{"Task2", false},
				},
				groups: []*domain.Group{},
			},
			want: 0.0,
		},
		"0%_NoTaskAndGroup": {
			in: input{
				tasks:  []*domain.Task{},
				groups: []*domain.Group{},
			},
			want: 0.0,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			c := domain.NewTodoItemContainer()
			for _, task := range tt.in.tasks {
				c.AddTask(task)
			}
			for _, group := range tt.in.groups {
				c.AddGroup(group)
			}
			got := calculateProgress(c)
			if got != tt.want {
				t.Errorf("calculateProgress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateTaskProgress(t *testing.T) {
	tests := map[string]struct {
		in   []*domain.Task
		want float64
	}{
		"100%": {
			[]*domain.Task{{"T1", true}, {"T2", true}},
			1.0,
		},
		"50%": {
			in:   []*domain.Task{{"T1", true}, {"T2", false}},
			want: 0.5,
		},
		"0%": {
			[]*domain.Task{{"T1", false}, {"T2", false}},
			0.0,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := calculateTaskProgress(tt.in)
			if got != tt.want {
				t.Errorf("calculateDoneTaskNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateDoneTaskNum(t *testing.T) {
	tests := map[string]struct {
		in   []*domain.Task
		want int
	}{
		"0DoneTask": {
			in: []*domain.Task{
				{"T1", false},
				{"T2", false},
				{"T3", false},
			},
			want: 0,
		},
		"1DoneTask": {
			in: []*domain.Task{
				{"T1", true},
				{"T2", false},
				{"T3", false},
			},
			want: 1,
		},
		"3DoneTasks": {
			in: []*domain.Task{
				{"T1", true},
				{"T2", true},
				{"T3", true},
				{"T4", false},
				{"T5", false},
				{"T6", false},
			},
			want: 3,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := calculateDoneTaskNum(tt.in)
			if got != tt.want {
				t.Errorf("calculateDoneTaskNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
