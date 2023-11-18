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
	upperPrefix := strings.ToUpper(prefix)
	lowerPrefix := strings.ToLower(prefix)

	if strings.HasPrefix(f.Line, upperPrefix) {
		f.Line = strings.TrimPrefix(f.Line, upperPrefix)
	} else if strings.HasPrefix(f.Line, lowerPrefix) {
		f.Line = strings.TrimPrefix(f.Line, lowerPrefix)
	}

	return f
}

func (f *LineFormatter) TrimSpace() {
	f.Line = strings.TrimSpace(f.Line)
}

func (f *LineFormatter) GetStatusString() string {
	if f.HasPrefix("X") || f.HasPrefix("x") {
		return "X"
	}
	return " "
}

func IsGroupTitle(s string) bool {
	if !strings.HasPrefix(s, "-") {
		return false
	}
	s = strings.TrimPrefix(s, "-")
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "[") {
		return false
	}
	return true
}

func IsGroupTaskString(s string) bool {
	if !strings.HasPrefix(s, " ") {
		return false
	}
	formatter := NewLineFormatter(strings.TrimSpace(s))

	if !formatter.HasPrefix("-") {
		return false
	}
	formatter.TrimPrefix("-").TrimSpace()

	if !formatter.HasPrefix("[") {
		return false
	}
	formatter.TrimPrefix("[").TrimSpace()

	if formatter.HasPrefix("X") || formatter.HasPrefix("x") {
		formatter.TrimPrefix("X").TrimSpace()
	}

	if !strings.HasPrefix(s, "]") {
		return false
	}

	return true
}

func GetGroupTitle(s string) string {
	s = strings.TrimPrefix(s, "-")
	return strings.TrimSpace(s)
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

func FormatGroupTitleString(s string) string {
	if !IsGroupTitle(s) {
		return ""
	}
	title := GetGroupTitle(s)
	return fmt.Sprintf("- %s", title)
}

func FormatGroupTaskString(s string) string {
	if !strings.HasPrefix(s, " ") {
		return ""
	}
	noSpaceStr := strings.TrimSpace(s)
	formattedString := FormatTaskString(noSpaceStr)
	if formattedString == "" {
		return ""
	}
	return fmt.Sprintf("  %s", formattedString)
}

func FormatTaskString(s string) string {
	if !strings.HasPrefix(s, "-") {
		return ""
	}
	if IsGroupTitle(s) {
		return fmt.Sprintf("- %s", GetGroupTitle(s))
	}

	formatter := NewLineFormatter(s)

	formatter.TrimPrefix("-").TrimSpace()

	if !formatter.HasPrefix("[") {
		return ""
	}
	formatter.TrimPrefix("[").TrimSpace()

	statusStr := formatter.GetStatusString()
	formatter.TrimPrefix(statusStr).TrimSpace()

	if !formatter.HasPrefix("]") {
		return ""
	}

	formatter.TrimPrefix("]").TrimSpace()

	return fmt.Sprintf("- [%s] %s", statusStr, formatter.Line)
}

func FormatGroupTitleString(s string) string {
	if !IsGroupTitle(s) {
		s = ""
	}
	title := GetGroupTitle(s)
	return fmt.Sprintf("- %s", title)
}

func FormatInGroupString(s string) string {
	if !strings.HasPrefix(s, " ") {
		return ""
	}
	noSpaceStr := strings.TrimSpace(s)
	formattedString := FormatTaskString(noSpaceStr)
	if formattedString == "" {
		return ""
	}
	return fmt.Sprintf("  %s", formattedString)
}
