package logic

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
			got, _ := FormatTaskStrings(tt.in)
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
			got, _ := FormatTaskStrings(tt.in)
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
			got, _ := FormatTaskStrings(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatTaskStrings() = %s, want %s", joinWithComma(got), joinWithComma(tt.want))
			}
		})
	}
}

func Test_formatTaskString(t *testing.T) {
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
			got, err := formatTaskString(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatTaskString() error = %v, wantErr = %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("FormatTaskString() = %q, want %q", got, tt.want)
			}
		})
	}
}

func Test_formatGroupTaskString(t *testing.T) {
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
			got, err := formatGroupTaskString(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatGroupTaskString() error = %v, wantErr = %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("FormatGroupTaskString() = %q, want %q", got, tt.want)
			}
		})
	}
}

func Test_formatGroupTitleString(t *testing.T) {
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
			got, err := formatGroupTitleString(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatGroupTitleString() error = %v, wantErr = %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("FormatGroupTitleString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getStatusString(t *testing.T) {
	tests := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		// Success cases
		"ValidTaskStringGoodFormat_Done":   {"- [ ] TaskName", " ", false},
		"ValidTaskStringGoodFormat_Undone": {"- [X] TaskName", "X", false},
		"ValidTaskStringBadFormat_Done":    {"-[]TaskName", " ", false},
		"ValidTaskStringBadFormat_Undone":  {"-[X]TaskName", "X", false},
		// Error cases
		"InValidTaskString_NoDash":         {"[ ] TaskName", "", true},
		"InValidTaskString_NoBracketStart": {"- ] TaskName", "", true},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := getStatusString(tt.in)
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

func Test_getGroupTitle(t *testing.T) {
	tests := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		// Success cases
		"ValidGroupStringGoodFormat": {"- GroupTitle", "GroupTitle", false},
		"ValidGroupStringBadFormat":  {"-GroupTitle", "GroupTitle", false},
		// Error cases
		"InvalidGroupString_NoDash": {"GroupTitle", "", true},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := getGroupTitle(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroupTitle() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("GetGroupTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isGroupTitle(t *testing.T) {
	tests := map[string]struct {
		in   string
		want bool
	}{
		"ValidGroupTitleString":   {"- GroupTitle", true},
		"InvalidGroupTitleString": {"GroupTitle", false},
		"TaskString":              {"- [ ] TaskName", false},
		"GroupTaskString":         {"  - [ ] TaskName", false},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := isGroupTitle(tt.in)
			if got != tt.want {
				t.Errorf("IsGroupTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isGroupTaskString(t *testing.T) {
	tests := map[string]struct {
		in   string
		want bool
	}{
		"ValidGroupTaskString":                  {"  - [ ] TaskName", true},
		"InvalidGroupTaskString_NoDash":         {"  [ ] TaskName", false},
		"InvalidGroupTaskString_NoBracketStart": {"  -  ] TaskName", false},
		"InvalidGroupTaskString_NoBracketEnd":   {"  - [  TaskName", false},
		"SingleTaskString":                      {"- [ ] TaskName", false},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := isGroupTaskString(tt.in)
			if got != tt.want {
				t.Errorf("IsGroupTaskString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSingleTaskString(t *testing.T) {
	tests := map[string]struct {
		in   string
		want bool
	}{
		"SingleTask_Done":               {"- [ ] I am SingleTask!", true},
		"SingleTask_Undone":             {"- [X] I am SingleTask!", true},
		"GroupTask_Done":                {"  - [X] I am SingleTask!", false},
		"GroupTask_Undone":              {"  - [X] I am SingleTask!", false},
		"Invalid_Undone_NoBracketStart": {"- X] I am SingleTask!", false},
		"Invalid_Undone_NoBracketEnd":   {"- [X I am SingleTask!", false},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := isSingleTaskString(tt.in); got != tt.want {
				t.Errorf("IsSingleTaskString() = %v, want %v", got, tt.want)
			}
		})
	}
}
