package code

import (
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

	f := NewLineFormatter(taskString)

	f.TrimPrefix("-").TrimSpace()

	if !f.HasPrefix("[") {
		return "", NoBracketStartError
	}
	f.TrimPrefix("[").TrimSpace()

	if f.HasPrefix("X") || f.HasPrefix("x") {
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
	f := NewLineFormatter(s)
	f.TrimPrefix("-").TrimSpace()
	if f.HasPrefix("[") {
		return false
	}
	return true
}

func IsGroupTaskString(s string) bool {
	if !strings.HasPrefix(s, " ") {
		return false
	}
	f := NewLineFormatter(strings.TrimSpace(s))

	if !f.HasPrefix("-") {
		return false
	}
	f.TrimPrefix("-").TrimSpace()

	if !f.HasPrefix("[") {
		return false
	}
	f.TrimPrefix("[").TrimSpace()

	if f.HasPrefix("X") || f.HasPrefix("x") {
		f.TrimPrefix("X").TrimSpace()
	}

	if !f.HasPrefix("]") {
		return false
	}

	return true
}

func IsSingleTaskString(s string) bool {
	if !strings.HasPrefix(s, "-") {
		return false
	}

	f := NewLineFormatter(s)
	f.TrimPrefix("-").TrimSpace()

	if !f.HasPrefix("[") {
		return false
	}
	f.TrimPrefix("[").TrimSpace()

	if f.HasPrefix("X") || f.HasPrefix("x") {
		f.TrimPrefix("X").TrimSpace()
	}

	if !f.HasPrefix("]") {
		return false
	}

	return true
}
