package line

import (
	"reflect"
	"testing"
)

func TestNewLineFormatter(t *testing.T) {
	tests := map[string]struct {
		line string
		want Line
	}{
		"Success": {line: "AAA-BBB", want: Line("AAA-BBB")},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := New(tt.line)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineFormatter_HasPrefix(t *testing.T) {
	tests := map[string]struct {
		line string
		in   string
		want bool
	}{
		"HasPrefix_True":  {line: "AAA-BBB", in: "AAA", want: true},
		"HasPrefix_False": {line: "AAA-BBB", in: "BBB", want: false},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			line := Line(tt.line)

			got := line.HasPrefix(tt.in)
			if got != tt.want {
				t.Errorf("HasPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineFormatter_TrimPrefix(t *testing.T) {
	tests := map[string]struct {
		line string
		in   string
		want Line
	}{
		"TrimPrefix_ExistPrefix":    {line: "AAA-BBB", in: "AAA", want: "-BBB"},
		"TrimPrefix_NotExistPrefix": {line: "AAA-BBB", in: "CCC", want: "AAA-BBB"},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			line := Line(tt.line)
			got := line.TrimPrefix(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrimPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineFormatter_TrimSpace(t *testing.T) {
	tests := map[string]struct {
		line string
		want Line
	}{
		"TrimSpace_Space":          {line: "  AAA-BBB  ", want: "AAA-BBB"},
		"TrimSpace_EscapeSequence": {line: "\nAAA-BBB\n", want: "AAA-BBB"},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			line := Line(tt.line)
			got := line.TrimSpace()
			if got != tt.want {
				t.Errorf("TrimSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
