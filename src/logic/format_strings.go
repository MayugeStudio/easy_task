package logic

import (
	"easy_task/src/logic/line"
	"errors"
	"fmt"
	"strings"
)

var (
	SyntaxError         = errors.New("format error")
	NoDashError         = fmt.Errorf("%w: no dash", SyntaxError)
	NoBracketStartError = fmt.Errorf("%w: no bracket start", SyntaxError)
	NoBracketEndError   = fmt.Errorf("%w: no bracket end", SyntaxError)
	InvalidIndentError  = fmt.Errorf("%w: no valid indent", SyntaxError)
)

func FormatTaskStrings(taskStrings []string) []string {
	result := make([]string, 0)
	errs := make([]error, 0)
	inGroup := false
	for _, str := range taskStrings {
		var formattedString string
		var err error
		if isGroupTitle(str) {
			formattedString, err = formatGroupTitleString(str)
			inGroup = true
		} else if inGroup && isGroupTaskString(str) {
			formattedString, err = formatGroupTaskString(str)
		} else if isSingleTaskString(str) {
			formattedString, err = formatTaskString(str)
			inGroup = false
		} else {
			if !strings.HasPrefix(str, "  ") {
				inGroup = false
			}
			continue
		}
		if err != nil {
			errs = append(errs, err)
			continue
		}
		result = append(result, formattedString)
	}
	return result
}

func formatGroupTitleString(s string) (string, error) {
	title, err := getGroupTitle(s)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("- %s", title), nil
}

func formatGroupTaskString(s string) (string, error) {
	if !strings.HasPrefix(s, " ") {
		return "", InvalidIndentError
	}
	noSpaceStr := strings.TrimSpace(s)
	formattedString, err := formatTaskString(noSpaceStr)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("  %s", formattedString), nil
}

func formatTaskString(s string) (string, error) {
	if !strings.HasPrefix(s, "-") {
		return "", NoDashError
	}

	l := line.New(s)

	l = l.TrimPrefix("-").TrimSpace()

	if !l.HasPrefix("[") {
		return "", NoBracketStartError
	}
	l = l.TrimPrefix("[").TrimSpace()

	statusStr, err := getStatusString(s)
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

func getStatusString(taskString string) (string, error) {
	if !strings.HasPrefix(taskString, "-") {
		return "", NoDashError
	}

	l := line.New(taskString)

	l = l.TrimPrefix("-").TrimSpace()

	if !l.HasPrefix("[") {
		return "", NoBracketStartError
	}
	l = l.TrimPrefix("[").TrimSpace()

	if l.HasPrefix("X") || l.HasPrefix("x") {
		return "X", nil
	}
	return " ", nil
}

func getGroupTitle(s string) (string, error) {
	if !strings.HasPrefix(s, "-") {
		return "", fmt.Errorf("%w: invalid group title %q", SyntaxError, s)
	}
	s = strings.TrimPrefix(s, "-")
	return strings.TrimSpace(s), nil
}

func isGroupTitle(s string) bool {
	if !strings.HasPrefix(s, "-") {
		return false
	}
	l := line.New(s)
	l = l.TrimPrefix("-").TrimSpace()
	if l.HasPrefix("[") {
		return false
	}
	return true
}

func isGroupTaskString(s string) bool {
	if !strings.HasPrefix(s, " ") {
		return false
	}
	l := line.New(strings.TrimSpace(s))

	if !l.HasPrefix("-") {
		return false
	}
	l = l.TrimPrefix("-").TrimSpace()

	if !l.HasPrefix("[") {
		return false
	}
	l = l.TrimPrefix("[").TrimSpace()

	if l.HasPrefix("X") || l.HasPrefix("x") {
		l = l.TrimPrefix("X").TrimSpace()
	}

	if !l.HasPrefix("]") {
		return false
	}

	return true
}

func isSingleTaskString(s string) bool {
	if !strings.HasPrefix(s, "-") {
		return false
	}

	l := line.New(s)
	l = l.TrimPrefix("-").TrimSpace()

	if !l.HasPrefix("[") {
		return false
	}
	l = l.TrimPrefix("[").TrimSpace()

	if l.HasPrefix("X") || l.HasPrefix("x") {
		l = l.TrimPrefix("X").TrimSpace()
	}

	if !l.HasPrefix("]") {
		return false
	}

	return true
}
