package code

import (
	"bytes"
	"testing"
)

func TestPrintTaskProgress(t *testing.T) {
	tests := map[string]struct {
		in      []*Task
		wantW   string
		wantErr bool
	}{
		"100%": {
			[]*Task{
				{"Task1", true},
				{"Task2", true},
			},
			"[########################################]100%",
			false,
		},
		"50%": {
			[]*Task{
				{"Task1", true},
				{"Task2", false},
			},
			"[####################                    ]50%",
			false,
		},
		"25%": {
			[]*Task{
				{"Task1", true},
				{"Task2", false},
				{"Task3", false},
				{"Task4", false},
			},
			"[##########                              ]25%",
			false,
		},
		"0%": {
			[]*Task{
				{"Task1", false},
				{"Task2", false},
			},
			"[                                        ]0%",
			false,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			w := &bytes.Buffer{}
			err := PrintTaskProgress(w, tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrintTaskProgress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("PrintTaskProgress() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
