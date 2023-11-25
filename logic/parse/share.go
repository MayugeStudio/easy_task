package parse

import (
	"strings"

	"github.com/MayugeStudio/easy_task/utils"
)

func isGroupTitle(s string) bool {
	if !strings.HasPrefix(s, "-") {
		return false
	}
	l := utils.NewLine(s)
	l = l.TrimPrefix("-").TrimSpace()
	return !l.HasPrefix("[")
}

func isGroupTaskString(s string) bool {
	if !strings.HasPrefix(s, " ") {
		return false
	}
	l := utils.NewLine(strings.TrimSpace(s))

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

	l := utils.NewLine(s)
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
