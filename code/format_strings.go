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

func (f *LineFormatter) Replace(old string, new string, n int) *LineFormatter {
	f.Line = strings.Replace(f.Line, old, new, n)
	return f
}

func (f *LineFormatter) TrimSpace() {
	f.Line = strings.TrimSpace(f.Line)
}

func GetStatusString(taskString string) string {
	if !strings.HasPrefix(taskString, "-") {
		return ""
	}

	f := NewLineFormatter(taskString)

	f.TrimPrefix("-").TrimSpace()

	if !f.HasPrefix("[") {
		return ""
	}
	f.TrimPrefix("[").TrimSpace()

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
		if IsGroupTitle(line) {
			groupTitle := FormatGroupTitleString(line)
			result = append(result, groupTitle)
			continue
		}

		if IsGroupTaskString(line) {
			formattedString := FormatGroupTaskString(line)
			if formattedString == "" {
				continue
			}
			result = append(result, formattedString)
			continue
		}

		formattedString := FormatTaskString(line) // TODO: Create FormatSingleTaskString function to improve readability.
		if formattedString == "" {
			continue
		}
		result = append(result, formattedString)
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

	formatter := NewLineFormatter(s)

	formatter.TrimPrefix("-").TrimSpace()

	if !formatter.HasPrefix("[") {
		return ""
	}
	formatter.TrimPrefix("[").TrimSpace()

	statusStr := GetStatusString(s)
	formatter.TrimPrefix(statusStr).TrimSpace()

	if !formatter.HasPrefix("]") {
		return ""
	}

	formatter.TrimPrefix("]").TrimSpace()

	return fmt.Sprintf("- [%s] %s", statusStr, formatter.Line)
}
