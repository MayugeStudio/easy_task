package format

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

func TestToValidStrings_OnlySingleTasks(t *testing.T) {
	tests := map[string]struct {
		in   []string
		want []string
	}{
		"ToValidStrings": {
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
			got, _ := ToValidStrings(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatTaskStrings() = %s, want %s", joinWithComma(got), joinWithComma(tt.want))
			}
		})
	}
}

func TestToValidStrings_OnlyGroupTasks(t *testing.T) {
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
			got, _ := ToValidStrings(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatTaskStrings() = %s, want %s", joinWithComma(got), joinWithComma(tt.want))
			}
		})
	}
}

func TestToValidStrings_MultiGroup(t *testing.T) {
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
			got, _ := ToValidStrings(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatTaskStrings() = %s, want %s", joinWithComma(got), joinWithComma(tt.want))
			}
		})
	}
}

func Test_toFormattedTaskString(t *testing.T) {
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

func Test_toFormattedGroupTaskString(t *testing.T) {
	tests := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		"Valid":         {in: "  - [ ] Buy the milk.", want: "  - [ ] Buy the milk."},
		"OneIndent":     {in: " - [ ] Buy the milk.", want: "  - [ ] Buy the milk."},
		"NoIndent":      {in: "- [ ] Buy the milk.", wantErr: true},
		"InvalidFormat": {in: "  - Buy the milk.", wantErr: true},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := toFormattedGroupTaskString(tt.in)
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

func Test_toFormattedGroupTitle(t *testing.T) {
	tests := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		"Valid":       {in: "- GroupTitle", want: "- GroupTitle"},
		"BadIndent":   {in: "-GroupTitle", want: "- GroupTitle"},
		"InvalidLine": {in: "GroupTitle", wantErr: true},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := toFormattedGroupTitle(tt.in)
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

func Test_toFormattedTaskStatus(t *testing.T) {
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

func Test_toFormattedModificationString(t *testing.T) {
	tests := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		"Valid_NoIndent":              {in: "<- [Tag]: Feature", want: "<- [Tag]: Feature"},
		"Valid_4Indent":               {in: "    <- [Tag]: Feature", want: "    <- [Tag]: Feature"},
		"Valid_BadFormat_NoIndent":    {in: "<-[Tag] :Feature", want: "<- [Tag]: Feature"},
		"Valid_BadFormat_4Indent":     {in: "    <-[Tag] :Feature", want: "    <- [Tag]: Feature"},
		"Invalid_InvalidModification": {in: "[Tag]: Feature", wantErr: true},
		"Invalid_NoBracketStart":      {in: "<- Tag]: Feature", wantErr: true},
		"Invalid_NoBracketEnd":        {in: "<- [Tag: Feature", wantErr: true},
		"Invalid_NoColon":             {in: "<- [Tag] Feature", wantErr: true},
		"Invalid_InvalidAttribute":    {in: "<- [tag]: Feature", wantErr: true},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := toFormattedModificationString(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("toFormattedModificationString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toFormattedModificationString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractGroupTitle(t *testing.T) {
	tests := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		"ValidGroupStringGoodFormat": {in: "- GroupTitle", want: "GroupTitle"},
		"ValidGroupStringBadFormat":  {in: "-GroupTitle", want: "GroupTitle"},
		"InvalidGroupString_NoDash":  {in: "GroupTitle", wantErr: true},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := extractGroupTitle(tt.in)
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
