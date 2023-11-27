package format

import (
	"reflect"
	"strings"
	"testing"
)

func TestToFormattedStrings(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in   []string
		want []string
	}{
		"BadFormatTaskStrings": {
			in: []string{
				"-[]Task1",
				"- [] Task2",
				"- []Task3",
			},
			want: []string{
				"- [ ] Task1",
				"- [ ] Task2",
				"- [ ] Task3",
			},
		},
		"BadFormatGroupStrings": {
			in: []string{
				"- Group",
				"  -[X] Task1",
				"  - [] Task2",
				"  - [ ]Task3",
			},
			want: []string{
				"- Group",
				"  - [X] Task1",
				"  - [ ] Task2",
				"  - [ ] Task3",
			},
		},
		"BadFormat2GroupStrings": {
			in: []string{
				"- Group1",
				"  -[X] Task1",
				"  -[]Task2",
				"  -[ ]Task3",
				"-Group2",
				"  -[X]Task1",
				"  - [] Task2",
			},
			want: []string{
				"- Group1",
				"  - [X] Task1",
				"  - [ ] Task2",
				"  - [ ] Task3",
				"- Group2",
				"  - [X] Task1",
				"  - [ ] Task2",
			},
		},
		"ContainsInvalidTaskString": {
			in: []string{
				"- [ ] Task1",
				"InvalidString",
				"- [ ] Task2",
			},
			want: []string{
				"- [ ] Task1",
				"- [ ] Task2",
			},
		},
		"ContainsInvalidIndentGroupTaskString": {
			in: []string{
				"- Group",
				"  - [X] Task1",
				"InvalidString.",
				"  - [ ] Task2",
			},
			want: []string{
				"- Group",
				"  - [X] Task1",
			},
		},
		"ContainsInvalidGroupTaskStringOtherThanInvalidIndent": {
			in: []string{
				"- Group",
				"  - [X] Task1",
				"  InvalidString",
				"  - [ ] Task2",
			},
			want: []string{
				"- Group",
				"  - [X] Task1",
				"  - [ ] Task2",
			},
		},
		"AllGroupTaskStringsAreInvalid": {
			in: []string{
				"- Group",
				"  Task1",
				"  Task2",
				"  Task3",
			},
			want: []string{
				"- Group",
			},
		},
		"InvalidGroupTitle": {
			in: []string{
				"Group",
				"  - [X] Task1",
				"  - [X] Task2",
				"  - [ ] Task3",
			},
			want: []string{},
		},
		"NestedGroup": {
			in: []string{
				"- Group1",
				"  - [X] Task1",
				"  - [ ] Task2",
				"  - Group2",
				"    - [ ] Task3",
				"    - [ ] Task4",
			},
			want: []string{
				"- Group1",
				"  - [X] Task1",
				"  - [ ] Task2",
				"  - Group2",
				"    - [ ] Task3",
				"    - [ ] Task4",
			},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, _ := ToFormattedStrings(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatTaskStrings() = \n%s \nwant \n%s", strings.Join(got, "\n"), strings.Join(tt.want, "\n"))
			}
		})
	}
}

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

func Test_toFormattedGroupTaskString(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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

func Test_toFormattedModificationString(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
