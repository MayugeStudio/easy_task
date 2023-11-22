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
