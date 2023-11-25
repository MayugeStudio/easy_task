package parse

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
			in: []string{
				"-[]Bake the bread.",
				"- [] Fry eggs.",
				"- []Prepare coffee.",
			},
			want: []string{
				"- [ ] Bake the bread.",
				"- [ ] Fry eggs.",
				"- [ ] Prepare coffee.",
			},
		},
		"ContainsInvalidTaskString": {
			in: []string{
				"- [ ] Bake the bread.",
				"Invalid TaskString.",
				"- [ ] Prepare coffee.",
			},
			want: []string{
				"- [ ] Bake the bread.",
				"- [ ] Prepare coffee.",
			},
		},
		"AllTaskStringsAreInvalid": {
			in: []string{
				"Bake the bread.",
				"Fry eggs.",
				"Prepare coffee.",
			},
			want: []string{},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, _ := formatTaskStrings(tt.in)
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
			in: []string{
				"- Eat breakfast.",
				// Child task must have Two indentations in the prefix.
				"  -[X]Bake the bread.",
				"  - [] Fry eggs.",
				"  - [ ]Prepare coffee.",
			},
			want: []string{
				"- Eat breakfast.",
				"  - [X] Bake the bread.",
				"  - [ ] Fry eggs.",
				"  - [ ] Prepare coffee.",
			},
		},
		"ContainsInvalidIndentChildTaskString": {
			in: []string{
				"- Eat breakfast.",
				"  - [X] Bake the bread.",
				"Invalid TaskString.",
				"  - [ ] Prepare coffee.",
			},
			want: []string{
				"- Eat breakfast.",
				"  - [X] Bake the bread.",
			},
		},
		"ContainsInvalidChildTaskStringOtherThanInvalidIndent": {
			in: []string{
				"- Eat breakfast.",
				"  - [X] Bake the bread.",
				"  Invalid TaskString.",
				"  - [ ] Prepare coffee.",
			},
			want: []string{
				"- Eat breakfast.",
				"  - [X] Bake the bread.",
				"  - [ ] Prepare coffee.",
			},
		},
		"AllTaskStringsAreInvalid": {
			in: []string{
				"- Eat breakfast.",
				"  Bake the bread.",
				"  Fry eggs.",
				"  Prepare coffee.",
			},
			want: []string{
				"- Eat breakfast.",
			},
		},
		"InvalidGroupTitleWithUndoneTasks": {
			in: []string{
				"Eat breakfast.",
				"  - [ ] Bake the bread.",
				"  - [ ] Fry eggs.",
				"  - [ ] Prepare coffee.",
			},
			want: []string{},
		},
		"InvalidGroupTitleWithDoneTasks": {
			in: []string{
				"Eat breakfast.",
				"  - [X] Bake the bread.",
				"  - [X] Fry eggs.",
				"  - [X] Prepare coffee.",
			},
			want: []string{},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, _ := formatTaskStrings(tt.in)
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
			in: []string{
				"- Eat breakfast.",
				"  -[X] Bake the bread.",
				"  -[]Fry eggs.",
				"  -[ ]Prepare coffee.",
				"-Study English.",
				"  -[X]Watch english TV show.",
				"  - []Memorize english words.",
			},
			want: []string{
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
			got, _ := formatTaskStrings(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatTaskStrings() = %s, want %s", joinWithComma(got), joinWithComma(tt.want))
			}
		})
	}
}

func Test_formatTaskString(t *testing.T) {
	tests := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		// Success cases
		"Undone_Valid":                 {in: "- [ ] Buy the milk.", want: "- [ ] Buy the milk."},
		"Undone_BadIndentStartBracket": {in: "-[] Buy the milk.", want: "- [ ] Buy the milk."},
		"Undone_BadIndentEndBracket":   {in: "- []Buy the milk.", want: "- [ ] Buy the milk."},
		"Done_Valid":                   {in: "- [X] Buy the milk.", want: "- [X] Buy the milk."},
		"Done_BadIndentStartEnd":       {in: "-[X]Buy the milk.", want: "- [X] Buy the milk."},
		"Done_Valid_Lower":             {in: "- [x] Buy the milk.", want: "- [X] Buy the milk."},
		"Done_NoSpaceInBracket_Lower":  {in: "- [x] Buy the milk.", want: "- [X] Buy the milk."},
		"Done_BadIndentStartEnd_Lower": {in: "-[x]Buy the milk.", want: "- [X] Buy the milk."},
		"Done_NoDash":                  {in: "[X] No Dash.", wantErr: true},
		// Error cases
		"Done_NoBracketStart": {in: "- X] No BracketStart.", wantErr: true},
		"Done_NoBracketEnd":   {in: "- [X No BracketEnd.", wantErr: true},
		"Done_NoDash_Lower":   {in: "[x] No Dash.", wantErr: true},
		"Undone_NoDash":       {in: "[ ] Buy the milk.", wantErr: true},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := formatTaskString(tt.in)
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

func Test_formatGroupTaskString(t *testing.T) {
	tests := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		// Success cases
		"Valid":     {in: "  - [ ] Buy the milk.", want: "  - [ ] Buy the milk."},
		"OneIndent": {in: " - [ ] Buy the milk.", want: "  - [ ] Buy the milk."},
		// Error cases
		"NoIndent":      {in: "- [ ] Buy the milk.", wantErr: true},
		"InvalidFormat": {in: "  - Buy the milk.", wantErr: true},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := formatGroupTaskString(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatGroupTaskString() error = %v, wantErr = %v", err, tt.wantErr)
				return
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
		"Valid":     {in: "- GroupTitle", want: "- GroupTitle"},
		"BadIndent": {in: "-GroupTitle", want: "- GroupTitle"},
		// Error cases
		"InvalidLine": {in: "GroupTitle", wantErr: true},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := formatGroupTitleString(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatGroupTitleString() error = %v, wantErr = %v", err, tt.wantErr)
				return
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
		"ValidTaskStringGoodFormat_Done":   {in: "- [ ] TaskName", want: " "},
		"ValidTaskStringGoodFormat_Undone": {in: "- [X] TaskName", want: "X"},
		"ValidTaskStringBadFormat_Done":    {in: "-[]TaskName", want: " "},
		"ValidTaskStringBadFormat_Undone":  {in: "-[X]TaskName", want: "X"},
		// Error cases
		"InValidTaskString_NoDash":         {in: "[ ] TaskName", wantErr: true},
		"InValidTaskString_NoBracketStart": {in: "- ] TaskName", wantErr: true},
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
		"ValidGroupStringGoodFormat": {in: "- GroupTitle", want: "GroupTitle"},
		"ValidGroupStringBadFormat":  {in: "-GroupTitle", want: "GroupTitle"},
		// Error cases
		"InvalidGroupString_NoDash": {in: "GroupTitle", wantErr: true},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := getGroupTitle(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroupTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
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
		"ValidGroupTitleString":   {in: "- GroupTitle", want: true},
		"InvalidGroupTitleString": {in: "GroupTitle", want: false},
		"TaskString":              {in: "- [ ] TaskName", want: false},
		"GroupTaskString":         {in: "  - [ ] TaskName", want: false},
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
		"ValidGroupTaskString":                  {in: "  - [ ] TaskName", want: true},
		"InvalidGroupTaskString_NoDash":         {in: "  [ ] TaskName", want: false},
		"InvalidGroupTaskString_NoBracketStart": {in: "  -  ] TaskName", want: false},
		"InvalidGroupTaskString_NoBracketEnd":   {in: "  - [  TaskName", want: false},
		"SingleTaskString":                      {in: "- [ ] TaskName", want: false},
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
		"SingleTask_Done":               {in: "- [ ] I am SingleTask!", want: true},
		"SingleTask_Undone":             {in: "- [X] I am SingleTask!", want: true},
		"GroupTask_Done":                {in: "  - [X] I am SingleTask!", want: false},
		"GroupTask_Undone":              {in: "  - [X] I am SingleTask!", want: false},
		"Invalid_Undone_NoBracketStart": {in: "- X] I am SingleTask!", want: false},
		"Invalid_Undone_NoBracketEnd":   {in: "- [X I am SingleTask!", want: false},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := isSingleTaskString(tt.in)
			if got != tt.want {
				t.Errorf("IsSingleTaskString() = %v, want %v", got, tt.want)
			}
		})
	}
}
