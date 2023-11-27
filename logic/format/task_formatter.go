package format

import (
	"fmt"
	"strings"

	"github.com/MayugeStudio/easy_task/logic/internal/share"
	"github.com/MayugeStudio/easy_task/utils"
)

func toFormattedTaskString(s string) (string, error) {
	indentLevel := share.GetIndentLevel(s)
	indentStr := strings.Repeat(" ", indentLevel)
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

	return fmt.Sprintf("%s- [%s] %s", indentStr, statusStr, l), nil
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
	// FIXME: Other strings are ignored.
	return " ", nil
}
