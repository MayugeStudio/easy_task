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

func TestFormatTaskString_Task(t *testing.T) {
	validStringDone := "- [X] Buy the milk."
	validStringUndone := "- [ ] Buy the milk."
	errString := ""
	tests := map[string]struct {
		in   string
		want string
	}{
		"Undone_Valid":                 {"- [ ] Buy the milk.", validStringUndone},
		"Undone_NoDash":                {"[ ] Buy the milk.", errString},
		"Undone_BadIndentStartBracket": {"-[] Buy the milk.", validStringUndone},
		"Undone_BadIndentEndBracket":   {"- []Buy the milk.", validStringUndone},
		"Done_Valid":                   {"- [X] Buy the milk.", validStringDone},
		"Done_NoDash":                  {"[X] No Dash.", errString},
		"Done_NoSpaceInBracket":        {"- [X] Buy the milk.", validStringDone},
		"Done_BadIndentStartEnd":       {"-[X]Buy the milk.", validStringDone},
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

func TestFormatTaskString_Group(t *testing.T) {
	validString := "- Buy the milk task group."
	errString := ""
	tests := map[string]struct {
		in   string
		want string
	}{
		"Valid":     {"- Buy the milk task group.", validString},
		"BadIndent": {"-Buy the milk task group.", validString},
		"NoDash":    {"Buy the milk task group.", errString},
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
