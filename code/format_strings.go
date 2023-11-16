package code

import "strings"

type lineS struct {
	line string
}

func (l *lineS) HasPrefix(prefix string) bool {
	return strings.HasPrefix(l.line, prefix)
}

func (l *lineS) TrimPrefix(prefix string) {
	l.line = strings.TrimPrefix(l.line, prefix)
}

func (l *lineS) TrimSpace() {
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

	if !ls.HasPrefix("-") {
		return ""
	}

	ls.TrimPrefix("-")
	b.WriteString("-")
	b.WriteString(" ")
	ls.TrimSpace()

	if ls.HasPrefix("[") {
		b.WriteString("[")
		ls.TrimPrefix("[")
	} else {
		// task group string
		b.WriteString(ls.line)
		return b.String()
	}
	ls.TrimSpace()

	if ls.HasPrefix("X") {
		b.WriteString("X")
		ls.TrimPrefix("X")
	} else {
		b.WriteString(" ")
	}
	ls.TrimSpace()

	if ls.HasPrefix("]") {
		b.WriteString("]")
		ls.TrimPrefix("]")
	} else {
		return ""
	}
	b.WriteString(" ")
	ls.TrimSpace()

	b.WriteString(ls.line)
	return b.String()
}
