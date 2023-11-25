package print

import (
	"bytes"
	"testing"

	"github.com/MayugeStudio/easy_task/domain"
)

func TestPrintProgress(t *testing.T) {
	tests := map[string]struct {
		in      []*domain.Task
		wantW   string
		wantErr bool
	}{
		"100%": {
			in:      []*domain.Task{{"T1", true}},
			wantW:   "[########################################]100.0%",
			wantErr: false,
		},
		"50%": {
			in:      []*domain.Task{{"T1", true}, {"T2", false}},
			wantW:   "[####################                    ]50.0%",
			wantErr: false,
		},
		"0%": {
			in:      []*domain.Task{{"T1", false}},
			wantW:   "[                                        ]0.0%",
			wantErr: false,
		},
		"NonTask": {
			in:      []*domain.Task{},
			wantW:   "[                                        ]0.0%",
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
			want: "[########################################]100.0%",
		},
		"50%": {
			in:   input{progress: 0.5, length: 40},
			want: "[####################                    ]50.0%",
		},
		"0%": {
			in:   input{progress: 0, length: 40},
			want: "[                                        ]0.0%",
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
