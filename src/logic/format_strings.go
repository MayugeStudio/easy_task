package logic

import (
	"easy_task/src/logic/line"
	"errors"
	"fmt"
	"strings"
)

func FormatTaskStrings(taskStrings []string) []string {
	result := make([]string, 0)
	errs := make([]error, 0)
	inGroup := false
	for _, str := range taskStrings {
		var formattedString string
		var err error
		if IsGroupTitle(str) {
			formattedString, err = FormatGroupTitleString(str)
			inGroup = true
		} else if inGroup && IsGroupTaskString(str) {
			formattedString, err = FormatGroupTaskString(str)
		} else if IsSingleTaskString(str) {
			formattedString, err = FormatTaskString(str)
			inGroup = false
		} else {
			if !strings.HasPrefix(str, "  ") {
				inGroup = false
			}
			continue
		}
		if err != nil {
			errs = append(errs, err)
			if errors.Is(err, InvalidIndentError) {
				inGroup = false
			}
			continue
		}
		result = append(result, formattedString)
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
		return "", InvalidIndentError
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

	l := line.New(s)

	l = l.TrimPrefix("-").TrimSpace()

	if !l.HasPrefix("[") {
		return "", NoBracketStartError
	}
	l = l.TrimPrefix("[").TrimSpace()

	statusStr, err := GetStatusString(s)
	if err != nil {
		return "", err
	}
	l = l.TrimPrefix(statusStr).TrimSpace()

	if !l.HasPrefix("]") {
		return "", NoBracketEndError
	}

	l = l.TrimPrefix("]").TrimSpace()

	return fmt.Sprintf("- [%s] %s", statusStr, l), nil
}
