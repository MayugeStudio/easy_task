package code

import (
	"testing"
)

func Test_formatLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := map[string]struct {
		args args
		want string
	}{
		"TaskLine_Valid":            {args{line: "- [ ] Buy the milk."}, "- [ ] Buy the milk."},
		"UndoneTaskLine_BadIndent1": {args{line: "-[] Buy the milk."}, "- [ ] Buy the milk."},
		"UndoneTaskLine_BadIndent2": {args{line: "- []Buy the milk."}, "- [ ] Buy the milk."},
		"UndoneTaskLine_BadIndent3": {args{line: "- []Buy the milk."}, "- [ ] Buy the milk."},
		"DoneTaskLine_BadIndent1":   {args{line: "-[X] Buy the milk."}, "- [X] Buy the milk."},
		"DoneTaskLine_BadIndent2":   {args{line: "- [X]Buy the milk."}, "- [X] Buy the milk."},
		"DoneTaskLine_BadIndent3":   {args{line: "-[X]Buy the milk."}, "- [X] Buy the milk."},
		"NotStartWithDash":          {args{line: "[] notStartWithDash."}, ""},
		"InvalidString":             {args{line: "-A[] 'A' is invalid."}, ""},
		"GroupLine_Valid":           {args{line: "-Buy the milk task group."}, "- Buy the milk task group."},
		"GroupLine_BadIndent":       {args{line: "-Buy the milk task group."}, "- Buy the milk task group."},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := FormatTaskString(tt.args.line)
			if got != tt.want {
				t.Errorf("FormatTaskString() = %q, want %q", got, tt.want)
			}
		})
	}
}
