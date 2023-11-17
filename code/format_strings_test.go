package code

import (
	"reflect"
	"testing"
)

func TestFormatTaskStrings(t *testing.T) {
	tests := map[string]struct {
		in   []string
		want []string
	}{
		"TaskStrings": {
			[]string{"-[]Bake the bread.", "- [] Fry eggs.", "- []Prepare coffee."},
			[]string{"- [ ] Bake the bread.", "- [ ] Fry eggs.", "- [ ] Prepare coffee."},
		},
		"ContainsInvalidTaskString": {
			[]string{"- [ ] Bake the bread.", "Invalid TaskString.", "- [ ] Prepare coffee."},
			[]string{"- [ ] Bake the bread.", "- [ ] Prepare coffee."},
		},
		"AllTaskStringsAreInvalid": {
			[]string{"Bake the bread.", "Fry eggs.", "Prepare coffee."},
			[]string{},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := FormatTaskStrings(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatTaskStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		"DoneTaskLine_LowerCase":    {"-[x]Buy the milk.", "- [X] Buy the milk."},
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
