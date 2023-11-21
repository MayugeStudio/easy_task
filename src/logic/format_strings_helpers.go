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

func GetStatusString(taskString string) (string, error) {
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

func GetGroupTitle(s string) (string, error) {
	if !strings.HasPrefix(s, "-") {
		return "", fmt.Errorf("%w: invalid group title %q", SyntaxError, s)
	}
	s = strings.TrimPrefix(s, "-")
	return strings.TrimSpace(s), nil
}

func IsGroupTitle(s string) bool {
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

func IsGroupTaskString(s string) bool {
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

func IsSingleTaskString(s string) bool {
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
