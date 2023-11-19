package code

import (
	"fmt"
	"strings"
)

func FormatTaskStrings(taskStrings []string) []string {
	result := make([]string, 0)
	errs := make([]error, 0)
	inGroup := false
	for _, line := range taskStrings {
		var formattedString string
		var err error
		if IsGroupTitle(line) {
			inGroup = true
			formattedString, err = FormatGroupTitleString(line)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			result = append(result, formattedString)
			continue
		} else if IsGroupTaskString(line) && inGroup {
			formattedString, err = FormatGroupTaskString(line)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			result = append(result, formattedString)
			continue
		} else {
			formattedString, err = FormatTaskString(line)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			result = append(result, formattedString)
			inGroup = false
		}
	}
	return result
}

func FormatGroupTitleString(s string) (string, error) {
	title, err := GetGroupTitle(s)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("- %s", title), nil
}

func FormatGroupTaskString(s string) (string, error) {
	if !strings.HasPrefix(s, " ") {
		return "", NoValidIndentError
	}
	noSpaceStr := strings.TrimSpace(s)
	formattedString, err := FormatTaskString(noSpaceStr)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("  %s", formattedString), nil
}

func FormatTaskString(s string) (string, error) {
	if !strings.HasPrefix(s, "-") {
		return "", NoDashError
	}

	formatter := NewLineFormatter(s)

	formatter.TrimPrefix("-").TrimSpace()

	if !formatter.HasPrefix("[") {
		return "", NoBracketStartError
	}
	formatter.TrimPrefix("[").TrimSpace()

	statusStr, err := GetStatusString(s)
	if err != nil {
		return "", err
	}
	formatter.TrimPrefix(statusStr).TrimSpace()

	if !formatter.HasPrefix("]") {
		return "", NoBracketEndError
	}

	formatter.TrimPrefix("]").TrimSpace()

	return fmt.Sprintf("- [%s] %s", statusStr, formatter.Line), nil
}
