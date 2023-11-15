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

func FormatLines(lines []string) []string {
	result := make([]string, 0)
	for _, line := range lines {
		if fl := formatLine(line); fl != "" {
			result = append(result, fl)
		} else {
			continue
		}
	}
	return result
}

func formatLine(line string) string {
	var b strings.Builder

	ls := lineS{line}

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
		return ""
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
