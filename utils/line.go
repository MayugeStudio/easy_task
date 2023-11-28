package utils

import "strings"

type Line string

func NewLine(l string) Line {
	return Line(l)
}

func (l Line) HasPrefix(prefix string) bool {
	return strings.HasPrefix(l.String(), prefix)
}

func (l Line) TrimPrefix(prefix string) Line {
	if len(prefix) == 1 {
		upperPrefix := strings.ToUpper(prefix)
		lowerPrefix := strings.ToLower(prefix)

		result := Line(strings.TrimPrefix(l.String(), upperPrefix))
		result = Line(strings.TrimPrefix(result.String(), lowerPrefix))
		return result
	} else {
		return Line(strings.TrimPrefix(l.String(), prefix))
	}
}

func (l Line) TrimSpace() Line {
	return Line(strings.TrimSpace(l.String()))
}

func (l Line) Replace(old, new string, n int) Line {
	return Line(strings.Replace(l.String(), old, new, n))
}

func (l Line) String() string {
	return string(l)
}
