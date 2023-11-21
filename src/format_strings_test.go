package src

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func joinWithComma(elems []string) string {
	newElems := make([]string, 0)
	for _, elem := range elems {
		newElems = append(newElems, fmt.Sprintf("%q", elem))
	}
	return "[" + strings.Join(newElems, ", ") + "]"
}

func TestFormatTaskStrings_OnlySingleTasks(t *testing.T) {
	tests := map[string]struct {
		in   []string
		want []string
	}{
		"TaskStrings": {
			[]string{
				"-[]Bake the bread.",
				"- [] Fry eggs.",
				"- []Prepare coffee.",
			},
			[]string{
				"- [ ] Bake the bread.",
				"- [ ] Fry eggs.",
				"- [ ] Prepare coffee.",
			},
		},
		"ContainsInvalidTaskString": {
			[]string{
				"- [ ] Bake the bread.",
				"Invalid TaskString.",
				"- [ ] Prepare coffee.",
			},
			[]string{
				"- [ ] Bake the bread.",
				"- [ ] Prepare coffee.",
			},
		},
		"AllTaskStringsAreInvalid": {
			[]string{
				"Bake the bread.",
				"Fry eggs.",
				"Prepare coffee.",
			},
			[]string{},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := FormatTaskStrings(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatTaskStrings() = %s, want %s", joinWithComma(got), joinWithComma(tt.want))
			}
		})
	}
}

func TestFormatTaskStrings_OnlyGroupTasks(t *testing.T) {
	tests := map[string]struct {
		in   []string
		want []string
	}{
		"ValidGroupTaskString": {
			[]string{
				"- Eat breakfast.",
				// Child task must have Two indentations in the prefix.
				"  -[X]Bake the bread.",
				"  - [] Fry eggs.",
				"  - [ ]Prepare coffee.",
			},
			[]string{
				"- Eat breakfast.",
				"  - [X] Bake the bread.",
				"  - [ ] Fry eggs.",
				"  - [ ] Prepare coffee.",
			},
		},
		"ContainsInvalidIndentChildTaskString": {
			[]string{
				"- Eat breakfast.",
				"  - [X] Bake the bread.",
				"Invalid TaskString.",
				"  - [ ] Prepare coffee.",
			},
			[]string{
				"- Eat breakfast.",
				"  - [X] Bake the bread.",
			},
		},
		"ContainsInvalidChildTaskStringOtherThanInvalidIndent": {
			[]string{
				"- Eat breakfast.",
				"  - [X] Bake the bread.",
				"  Invalid TaskString.",
				"  - [ ] Prepare coffee.",
			},
			[]string{
				"- Eat breakfast.",
				"  - [X] Bake the bread.",
				"  - [ ] Prepare coffee.",
			},
		},
		"AllTaskStringsAreInvalid": {
			[]string{
				"- Eat breakfast.",
				"  Bake the bread.",
				"  Fry eggs.",
				"  Prepare coffee.",
			},
			[]string{
				"- Eat breakfast.",
			},
		},
		"InvalidGroupTitleWithUndoneTasks": {
			[]string{
				"Eat breakfast.",
				"  - [ ] Bake the bread.",
				"  - [ ] Fry eggs.",
				"  - [ ] Prepare coffee.",
			},
			[]string{},
		},
		"InvalidGroupTitleWithDoneTasks": {
			[]string{
				"Eat breakfast.",
				"  - [X] Bake the bread.",
				"  - [X] Fry eggs.",
				"  - [X] Prepare coffee.",
			},
			[]string{},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := FormatTaskStrings(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatTaskStrings() = %s, want %s", joinWithComma(got), joinWithComma(tt.want))
			}
		})
	}
}

func TestFormatTaskStrings_MultiGroup(t *testing.T) {
	tests := map[string]struct {
		in   []string
		want []string
	}{
		"ValidGroupTaskString": {
			[]string{
				"- Eat breakfast.",
				"  -[X] Bake the bread.",
				"  -[]Fry eggs.",
				"  -[ ]Prepare coffee.",
				"-Study English.",
				"  -[X]Watch english TV show.",
				"  - []Memorize english words.",
			},
			[]string{
				"- Eat breakfast.",
				"  - [X] Bake the bread.",
				"  - [ ] Fry eggs.",
				"  - [ ] Prepare coffee.",
				"- Study English.",
				"  - [X] Watch english TV show.",
				"  - [ ] Memorize english words.",
			},
		},
		"ContainInValidGroupTaskString": {
			[]string{
				"- Eat breakfast.",
				"  -[X] Bake the bread.",
				"  Fry eggs.",
				"  -[ ]Prepare coffee.",
				"-Study English.",
				"  -[X]Watch english TV show.",
				"Memorize english words.",
			},
			[]string{
				"- Eat breakfast.",
				"  - [X] Bake the bread.",
				"  - [ ] Prepare coffee.",
				"- Study English.",
				"  - [X] Watch english TV show.",
			},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := FormatTaskStrings(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatTaskStrings() = %s, want %s", joinWithComma(got), joinWithComma(tt.want))
			}
		})
	}
}


func TestFormatTaskString(t *testing.T) {
	validStringDone := "- [X] Buy the milk."
	validStringUndone := "- [ ] Buy the milk."
	tests := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		// Success cases
		"Undone_Valid":                 {"- [ ] Buy the milk.", validStringUndone, false},
		"Undone_BadIndentStartBracket": {"-[] Buy the milk.", validStringUndone, false},
		"Undone_BadIndentEndBracket":   {"- []Buy the milk.", validStringUndone, false},
		"Done_Valid":                   {"- [X] Buy the milk.", validStringDone, false},
		"Done_BadIndentStartEnd":       {"-[X]Buy the milk.", validStringDone, false},
		"Done_Valid_Lower":             {"- [x] Buy the milk.", validStringDone, false},
		"Done_NoSpaceInBracket_Lower":  {"- [x] Buy the milk.", validStringDone, false},
		"Done_BadIndentStartEnd_Lower": {"-[x]Buy the milk.", validStringDone, false},
		"Done_NoDash":                  {"[X] No Dash.", "", true},
		// Error cases
		"Done_NoBracketStart": {"- X] No BracketStart.", "", true},
		"Done_NoBracketEnd":   {"- [X No BracketEnd.", "", true},
		"Done_NoDash_Lower":   {"[x] No Dash.", "", true},
		"Undone_NoDash":       {"[ ] Buy the milk.", "", true},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := FormatTaskString(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatTaskString() error = %v, wantErr = %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("FormatTaskString() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestFormatGroupTaskString(t *testing.T) {
	validGroupString := "  - [ ] Buy the milk."
	tests := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		// Success cases
		"Valid":     {"  - [ ] Buy the milk.", validGroupString, false},
		"OneIndent": {" - [ ] Buy the milk.", validGroupString, false},
		// Error cases
		"NoIndent":      {"- [ ] Buy the milk.", "", true},
		"InvalidFormat": {"  - Buy the milk.", "", true},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := FormatGroupTaskString(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatGroupTaskString() error = %v, wantErr = %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("FormatGroupTaskString() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestFormatGroupTitleString(t *testing.T) {
	tests := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		// Success
		"Valid":     {"- GroupTitle", "- GroupTitle", false},
		"BadIndent": {"-GroupTitle", "- GroupTitle", false},
		// Error cases
		"InvalidLine": {"GroupTitle", "", true},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := FormatGroupTitleString(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatGroupTitleString() error = %v, wantErr = %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("FormatGroupTitleString() = %v, want %v", got, tt.want)
			}
		})
	}
}
