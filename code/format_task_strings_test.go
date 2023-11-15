package code

import (
	"testing"
)

func TestFormatTaskString(t *testing.T) {
	tests := map[string]struct {
		in   string
		want string
	}{
		"TaskLine_Valid":            {"- [ ] Buy the milk.", "- [ ] Buy the milk."},
		"UndoneTaskLine_BadIndent1": {"-[] Buy the milk.", "- [ ] Buy the milk."},
		"UndoneTaskLine_BadIndent2": {"- []Buy the milk.", "- [ ] Buy the milk."},
		"UndoneTaskLine_BadIndent3": {"- []Buy the milk.", "- [ ] Buy the milk."},
		"DoneTaskLine_BadIndent1":   {"-[X] Buy the milk.", "- [X] Buy the milk."},
		"DoneTaskLine_BadIndent2":   {"- [X]Buy the milk.", "- [X] Buy the milk."},
		"DoneTaskLine_BadIndent3":   {"-[X]Buy the milk.", "- [X] Buy the milk."},
		"NotStartWithDash":          {"[] notStartWithDash.", ""},
		"GroupLine_Valid":           {"-Buy the milk task group.", "- Buy the milk task group."},
		"GroupLine_BadIndent":       {"-Buy the milk task group.", "- Buy the milk task group."},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := FormatTaskString(tt.in)
			if got != tt.want {
				t.Errorf("FormatTaskString() = %q, want %q", got, tt.want)
			}
		})
	}
}
