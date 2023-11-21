package logic

import "strings"

type LineFormatter string

func NewLineFormatter(line string) LineFormatter {
	return LineFormatter(line)
}

func (l LineFormatter) HasPrefix(prefix string) bool {
	return strings.HasPrefix(l.toString(), prefix)
}

func (l LineFormatter) TrimPrefix(prefix string) LineFormatter {
	upperPrefix := strings.ToUpper(prefix)
	lowerPrefix := strings.ToLower(prefix)

	line := LineFormatter(strings.TrimPrefix(l.toString(), upperPrefix))
	line = LineFormatter(strings.TrimPrefix(line.toString(), lowerPrefix))
	return line
}

func (l LineFormatter) TrimSpace() LineFormatter {
	return LineFormatter(strings.TrimSpace(l.toString()))
}

func (l LineFormatter) toString() string {
	return string(l)
}
