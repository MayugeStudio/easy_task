package format

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MayugeStudio/easy_task/logic/internal/share"
	"github.com/MayugeStudio/easy_task/utils"
)

var (
	errSyntax         = errors.New("format error")
	errNoDash         = fmt.Errorf("%w: no dash", errSyntax)
	errNoBracketStart = fmt.Errorf("%w: no bracket start", errSyntax)
	errNoBracketEnd   = fmt.Errorf("%w: no bracket end", errSyntax)
	errInvalidIndent  = fmt.Errorf("%w: no valid indent", errSyntax)
)

func ToValidStrings(taskStrings []string) ([]string, []error) {
	result := make([]string, 0)
	errs := make([]error, 0)
	inGroup := false
	for _, str := range taskStrings {
		var formattedString string
		var err error
		if share.IsGroupTitle(str) {
			formattedString, err = toFormattedGroupTitle(str)
			inGroup = true
		} else if inGroup && share.IsGroupTaskString(str) {
			formattedString, err = toFormattedGroupTaskString(str)
		} else if share.IsSingleTaskString(str) {
			formattedString, err = toFormattedTaskString(str)
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
	return result, errs
}

func toFormattedGroupTitle(s string) (string, error) {
	title, err := extractGroupTitle(s)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("- %s", title), nil
}

func toFormattedGroupTaskString(s string) (string, error) {
	if !strings.HasPrefix(s, " ") {
		return "", errInvalidIndent
	}
	noSpaceStr := strings.TrimSpace(s)
	formattedString, err := toFormattedTaskString(noSpaceStr)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("  %s", formattedString), nil
}

func toFormattedTaskString(s string) (string, error) {
	if !strings.HasPrefix(s, "-") {
		return "", errNoDash
	}

	l := utils.NewLine(s)

	l = l.TrimPrefix("-").TrimSpace()

	if !l.HasPrefix("[") {
		return "", errNoBracketStart
	}
	l = l.TrimPrefix("[").TrimSpace()

	statusStr, err := toFormattedTaskStatus(s)
	if err != nil {
		return "", err
	}
	l = l.TrimPrefix(statusStr).TrimSpace()

	if !l.HasPrefix("]") {
		return "", errNoBracketEnd
	}

	l = l.TrimPrefix("]").TrimSpace()

	return fmt.Sprintf("- [%s] %s", statusStr, l), nil
}

func toFormattedTaskStatus(s string) (string, error) {
	if !strings.HasPrefix(s, "-") {
		return "", errNoDash
	}

	l := utils.NewLine(s)

	l = l.TrimPrefix("-").TrimSpace()

	if !l.HasPrefix("[") {
		return "", errNoBracketStart
	}
	l = l.TrimPrefix("[").TrimSpace()

	if l.HasPrefix("X") || l.HasPrefix("x") {
		return "X", nil
	}
	return " ", nil
}

func extractGroupTitle(s string) (string, error) {
	if !strings.HasPrefix(s, "-") {
		return "", fmt.Errorf("%w: invalid group title %q", errSyntax, s)
	}
	s = strings.TrimPrefix(s, "-")
	return strings.TrimSpace(s), nil
}
