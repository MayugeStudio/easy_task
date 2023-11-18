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

func IsGroupTitle(s string) bool {
	if !strings.HasPrefix(s, "-") {
		return false
	}
	s = strings.TrimPrefix(s, "-")
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "[") {
		return false
	}
	return true
}

func IsGroupTaskString(s string) bool {
	if !strings.HasPrefix(s, " ") {
		return false
	}
	formatter := NewLineFormatter(strings.TrimSpace(s))

	if !formatter.HasPrefix("-") {
		return false
	}
	formatter.TrimPrefix("-").TrimSpace()

	if !formatter.HasPrefix("[") {
		return false
	}
	formatter.TrimPrefix("[").TrimSpace()

	if formatter.HasPrefix("X") || formatter.HasPrefix("x") {
		formatter.TrimPrefix("X").TrimSpace()
	}

	if !strings.HasPrefix(s, "]") {
		return false
	}

	return true
}

func GetGroupTitle(s string) string {
	s = strings.TrimPrefix(s, "-")
	return strings.TrimSpace(s)
}
