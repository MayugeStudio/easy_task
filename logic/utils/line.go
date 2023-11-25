package utils

import "strings"

type Line string

func New(l string) Line {
	return Line(l)
}

func (l Line) HasPrefix(prefix string) bool {
	return strings.HasPrefix(l.toString(), prefix)
}

func (l Line) TrimPrefix(prefix string) Line {
	upperPrefix := strings.ToUpper(prefix)
	lowerPrefix := strings.ToLower(prefix)

	result := Line(strings.TrimPrefix(l.toString(), upperPrefix))
	result = Line(strings.TrimPrefix(result.toString(), lowerPrefix))
	return result
}

func (l Line) TrimSpace() Line {
	return Line(strings.TrimSpace(l.toString()))
}

func (l Line) toString() string {
	return string(l)
}
