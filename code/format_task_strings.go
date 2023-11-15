package code

import "strings"

type lineS struct {
	line string
}

func (l *lineS) hasPrefix(prefix string) bool {
	return strings.HasPrefix(l.line, prefix)
}

func (l *lineS) trimPrefix(prefix string) {
	l.line = strings.TrimPrefix(l.line, prefix)
}

func (l *lineS) trimSpace() {
	l.line = strings.TrimSpace(l.line)
}

func FormatTaskStrings(taskStrings []string) []string {
	result := make([]string, 0)
	for _, line := range taskStrings {
		if fl := FormatTaskString(line); fl != "" {
			result = append(result, fl)
		} else {
			continue
		}
	}
	return result
}

func FormatTaskString(taskString string) string {
	var b strings.Builder

	ls := lineS{taskString}

	if !ls.hasPrefix("-") {
		return ""
	}

	ls.trimPrefix("-")
	b.WriteString("-")
	b.WriteString(" ")
	ls.trimSpace()

	if ls.hasPrefix("[") {
		b.WriteString("[")
		ls.trimPrefix("[")
	} else {
		// task group string
		b.WriteString(ls.line)
		return b.String()
	}
	ls.trimSpace()

	if ls.hasPrefix("X") {
		b.WriteString("X")
		ls.trimPrefix("X")
	} else {
		b.WriteString(" ")
	}
	ls.trimSpace()

	if ls.hasPrefix("]") {
		b.WriteString("]")
		ls.trimPrefix("]")
	} else {
		return ""
	}
	b.WriteString(" ")
	ls.trimSpace()

	b.WriteString(ls.line)
	return b.String()
}
