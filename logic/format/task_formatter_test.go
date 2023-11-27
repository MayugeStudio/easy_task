package format

import "testing"

func Test_toFormattedTaskString(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		"Undone_Valid":                 {in: "- [ ] Buy the milk.", want: "- [ ] Buy the milk."},
		"Undone_BadIndentStartBracket": {in: "-[] Buy the milk.", want: "- [ ] Buy the milk."},
		"Undone_BadIndentEndBracket":   {in: "- []Buy the milk.", want: "- [ ] Buy the milk."},
		"Done_Valid":                   {in: "- [X] Buy the milk.", want: "- [X] Buy the milk."},
		"Done_BadIndentStartEnd":       {in: "-[X]Buy the milk.", want: "- [X] Buy the milk."},
		"Done_Valid_Lower":             {in: "- [x] Buy the milk.", want: "- [X] Buy the milk."},
		"Done_NoSpaceInBracket_Lower":  {in: "- [x] Buy the milk.", want: "- [X] Buy the milk."},
		"Done_BadIndentStartEnd_Lower": {in: "-[x]Buy the milk.", want: "- [X] Buy the milk."},
		"Done_NoDash":                  {in: "[X] No Dash.", wantErr: true},
		"Done_NoBracketStart":          {in: "- X] No BracketStart.", wantErr: true},
		"Done_NoBracketEnd":            {in: "- [X No BracketEnd.", wantErr: true},
		"Done_NoDash_Lower":            {in: "[x] No Dash.", wantErr: true},
		"Undone_NoDash":                {in: "[ ] Buy the milk.", wantErr: true},
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
		"ValidTaskStringGoodFormat_Done":   {in: "- [ ] TaskName", want: " "},
		"ValidTaskStringGoodFormat_Undone": {in: "- [X] TaskName", want: "X"},
		"ValidTaskStringBadFormat_Done":    {in: "-[]TaskName", want: " "},
		"ValidTaskStringBadFormat_Undone":  {in: "-[X]TaskName", want: "X"},
		"InValidTaskString_NoDash":         {in: "[ ] TaskName", wantErr: true}, //FIXME: InValid -> Invalid
		"InValidTaskString_NoBracketStart": {in: "- ] TaskName", wantErr: true},
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
