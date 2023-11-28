package format

import (
	"fmt"

	"github.com/MayugeStudio/easy_task/utils"
)

func toFormattedTaskString(s string) (string, error) {
	l := utils.NewLine(s)
	if !l.HasPrefix("-") {
		return "", errNoDash
	}
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
	l := utils.NewLine(s).TrimSpace()
	if !l.HasPrefix("-") {
		return "", errNoDash
	}

	l = l.TrimPrefix("-").TrimSpace()

	if !l.HasPrefix("[") {
		return "", errNoBracketStart
	}
	l = l.TrimPrefix("[").TrimSpace()

	if l.HasPrefix("X") || l.HasPrefix("x") {
		return "X", nil
	}
	// FIXME: Other strings are ignored.
	return " ", nil
}
