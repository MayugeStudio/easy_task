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
		"ContainsInvalidGroupTaskStringOtherThanInvalidIndent": {
			in: []string{
				"- Group",
				"  - [X] Task1",
				"InvalidString",
				"  - [ ] Task2",
				"  InvalidString",
				"  - [ ] Task3",
			},
			want: []string{
				"- Group",
				"  - [X] Task1",
				"  - [ ] Task2",
				"  - [ ] Task3",
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
