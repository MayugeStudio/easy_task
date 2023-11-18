package code

import "strings"

func GetStatusString(taskString string) string {
	if !strings.HasPrefix(taskString, "-") {
		return ""
	}

	f := NewLineFormatter(taskString)

	f.TrimPrefix("-").TrimSpace()

	if !f.HasPrefix("[") {
		return ""
	}
	f.TrimPrefix("[").TrimSpace()

	if f.HasPrefix("X") || f.HasPrefix("x") {
		return "X"
	}
	return " "
}

func GetGroupTitle(s string) string {
	if !strings.HasPrefix(s, "-") {
		return ""
	}
	s = strings.TrimPrefix(s, "-")
	return strings.TrimSpace(s)
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
