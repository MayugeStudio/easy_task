package code

import "strings"

type LineFormatter struct {
	Line    string
	Builder strings.Builder
}

func NewLineFormatter(line string) *LineFormatter {
	return &LineFormatter{Line: line}
}

func (f *LineFormatter) HasPrefix(prefix string) bool {
	return strings.HasPrefix(f.Line, prefix)
}

func (f *LineFormatter) TrimPrefix(prefix string) {
	f.Line = strings.TrimPrefix(f.Line, prefix)
}

func (f *LineFormatter) TrimSpace() {
	f.Line = strings.TrimSpace(f.Line)
}

func FormatTaskStrings(taskStrings []string) []string {
	result := make([]string, 0)
	for _, line := range taskStrings {
		if fl := FormatTaskString(line); fl != "" {
			result = append(result, fl)
		}
	}
	return result
}

func FormatTaskString(taskString string) string {
	var b strings.Builder

	lineFormatter := NewLineFormatter(taskString)

	if !lineFormatter.HasPrefix("-") {
		return ""
	}

	lineFormatter.TrimPrefix("-")
	b.WriteString("-")
	b.WriteString(" ")
	lineFormatter.TrimSpace()

	if lineFormatter.HasPrefix("[") {
		b.WriteString("[")
		lineFormatter.TrimPrefix("[")
	} else {
		// task group string
		b.WriteString(lineFormatter.Line)
		return b.String()
	}
	lineFormatter.TrimSpace()

	if lineFormatter.HasPrefix("X") {
		b.WriteString("X")
		lineFormatter.TrimPrefix("X")
	} else {
		b.WriteString(" ")
	}
	lineFormatter.TrimSpace()

	if lineFormatter.HasPrefix("]") {
		b.WriteString("]")
		lineFormatter.TrimPrefix("]")
	} else {
		return ""
	}
	b.WriteString(" ")
	lineFormatter.TrimSpace()

	b.WriteString(lineFormatter.Line)
	return b.String()
}
