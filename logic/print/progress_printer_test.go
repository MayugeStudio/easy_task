package print

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/MayugeStudio/easy_task/domain"
)

func TestProgress(t *testing.T) { // FIX-TEST: Add group pattern.
	tests := map[string]struct {
		in      []bool
		wantW   string
		wantErr bool
	}{
		"100%": {
			in:      []bool{true},
			wantW:   "[########################################]100.0%",
			wantErr: false,
		},
		"50%": {
			in:      []bool{true, false},
			wantW:   "[####################                    ]50.0%",
			wantErr: false,
		},
		"0%": {
			in:      []bool{false},
			wantW:   "[                                        ]0.0%",
			wantErr: false,
		},
		"NonTask": {
			in:      []bool{},
			wantW:   "[                                        ]0.0%",
			wantErr: false,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			w := &bytes.Buffer{}
			items := domain.NewItems()
			for i, b := range tt.in { // FIX-TEST: This code is redundant. Duplicate task title is ok.
				items = append(items, newTask(fmt.Sprintf("T%d", i), b))
			}
			err := Progress(w, items)
			if (err != nil) != tt.wantErr {
				t.Errorf("Progress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Progress() gotW = %v, want %v", gotW, tt.wantW)
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
