package code

import (
	"reflect"
	"testing"
)

func TestNewLineFormatter(t *testing.T) {
	tests := map[string]struct {
		line string
		want *LineFormatter
	}{
		"Success": {"AAA-BBB", &LineFormatter{"AAA-BBB"}},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewLineFormatter(tt.line)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLineFormatter() = %v, want %v", got, tt.want)
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
		"HasPrefix_True":  {"AAA-BBB", "AAA", true},
		"HasPrefix_False": {"AAA-BBB", "BBB", false},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			f := &LineFormatter{
				Line: tt.line,
			}
			got := f.HasPrefix(tt.in)
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
		want string
	}{
		"TrimPrefix_ExistPrefix":    {"AAA-BBB", "AAA", "-BBB"},
		"TrimPrefix_NotExistPrefix": {"AAA-BBB", "CCC", "AAA-BBB"},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			f := &LineFormatter{
				Line: tt.line,
			}
			got := f.TrimPrefix(tt.in).Line
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrimPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineFormatter_TrimSpace(t *testing.T) {
	tests := map[string]struct {
		line string
		want string
	}{
		"TrimSpace_Space":          {"  AAA-BBB  ", "AAA-BBB"},
		"TrimSpace_EscapeSequence": {"\nAAA-BBB\n", "AAA-BBB"},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			f := &LineFormatter{
				Line: tt.line,
			}
			f.TrimSpace()
			if f.Line != tt.want {
				t.Errorf("TrimSpace() = %v, want %v", f.Line, tt.want)
			}
		})
	}
}
