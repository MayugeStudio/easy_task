package format

import "testing"

func Test_toFormattedTaskString(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		"Undone_Valid":          {in: "- [ ] Task", want: "- [ ] Task"},
		"Done_Valid":            {in: "- [X] Task", want: "- [X] Task"},
		"Done_Valid_Lower":      {in: "- [x] Task", want: "- [X] Task"},
		"BadIndentStartBracket": {in: "-[] Task", want: "- [ ] Task"},
		"BadIndentEndBracket":   {in: "- []Task", want: "- [ ] Task"},
		"BadIndentStartEnd":     {in: "-[X]Task", want: "- [X] Task"},
		"NoDash":                {in: "[X] Task", wantErr: true},
		"NoBracketStart":        {in: "- X] Task", wantErr: true},
		"NoBracketEnd":          {in: "- [X Task", wantErr: true},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := toFormattedTaskString(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatTaskString() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FormatTaskString() = %q, want %q", got, tt.want)
			}
		})
	}
}

func Test_toFormattedTaskStatus(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		"GoodFormat_Done":         {in: "- [ ] Task", want: " "},
		"GoodFormat_Undone":       {in: "- [X] Task", want: "X"},
		"BadFormat_Done":          {in: "-[]Task", want: " "},
		"BadFormat_Undone":        {in: "-[X]Task", want: "X"},
		"BadFormat_InvalidStatus": {in: "-[A]Task", want: " "},
		"NoDash":                  {in: "[ ] Task", wantErr: true},
		"NoBracketStart":          {in: "- ] Task", wantErr: true},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := toFormattedTaskStatus(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStatusString() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetStatusString() = %v, want %v", got, tt.want)
			}
		})
	}
}
