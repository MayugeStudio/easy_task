package code

import (
	"fmt"
	"strings"
)

type LineFormatter struct {
	Line string
}

func NewLineFormatter(line string) *LineFormatter {
	return &LineFormatter{Line: line}
}

func (f *LineFormatter) HasPrefix(prefix string) bool {
	return strings.HasPrefix(f.Line, prefix)
}

func (f *LineFormatter) TrimPrefix(prefix string) *LineFormatter {
	f.Line = strings.TrimPrefix(f.Line, strings.ToUpper(prefix))
	f.Line = strings.TrimPrefix(f.Line, strings.ToLower(prefix))
	return f
}

func (f *LineFormatter) TrimSpace() {
	f.Line = strings.TrimSpace(f.Line)
}

func FormatTaskStrings(taskStrings []string) []string {
	result := make([]string, 0)
	for _, line := range taskStrings {
		if fl := FormatTaskString(line); fl != "" {
			result = append(result, fl)
		}
	}
	return result
}

func FormatTaskString(taskString string) string {
	if !strings.HasPrefix(taskString, "-") {
		return ""
	}
	if IsGroup(taskString) {
		return fmt.Sprintf("- %s", getGroupTitle(taskString))
	}

	formatter := NewLineFormatter(taskString)

	formatter.TrimPrefix("-").TrimSpace()

	if !formatter.HasPrefix("[") {
		return ""
	}
	formatter.TrimPrefix("[").TrimSpace()

	statusStr := GetStatusString(formatter)
	formatter.TrimPrefix(statusStr).TrimSpace()

	if !formatter.HasPrefix("]") {
		return ""
	}

	formatter.TrimPrefix("]").TrimSpace()

	return fmt.Sprintf("- [%s] %s", statusStr, formatter.Line)
}

func IsGroup(s string) bool {
	s = strings.TrimPrefix(s, "-")
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "[") {
		return false
	}
	return true
}

func getGroupTitle(s string) string {
	return strings.TrimSpace(strings.TrimPrefix(s, "-"))
}

func GetStatusString(f *LineFormatter) string {
	if f.HasPrefix("X") || f.HasPrefix("x") {
		return "X"
	}
	return " "
}
