package code

import (
	"fmt"
	"strings"
)

func FormatTaskStrings(taskStrings []string) []string {
	result := make([]string, 0)
	inGroup := false
	for _, line := range taskStrings {
		if IsGroupTitle(line) {
			inGroup = true
			groupTitle := FormatGroupTitleString(line)
			result = append(result, groupTitle)
			continue
		}

		if IsGroupTaskString(line) && inGroup {
			formattedString := FormatGroupTaskString(line)
			result = append(result, formattedString)
			continue
		}

		formattedString := FormatTaskString(line) // TODO: Create FormatSingleTaskString function to improve readability.
		if formattedString == "" {
			continue
		}
		result = append(result, formattedString)
		inGroup = false
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
