package code

import "strings"

type LineFormatter struct {
	Line string
	b    strings.Builder
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

func (f *LineFormatter) WriteString(s string) {
	f.b.WriteString(s)
}

func (f *LineFormatter) String() string {
	return f.b.String()
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
	lineFormatter := NewLineFormatter(taskString)

	if !lineFormatter.HasPrefix("-") {
		return ""
	}

	lineFormatter.TrimPrefix("-")
	lineFormatter.WriteString("-")
	lineFormatter.WriteString(" ")
	lineFormatter.TrimSpace()

	if lineFormatter.HasPrefix("[") {
		lineFormatter.WriteString("[")
		lineFormatter.TrimPrefix("[")
	} else {
		// task group string
		lineFormatter.WriteString(lineFormatter.Line)
		return lineFormatter.String()
	}
	lineFormatter.TrimSpace()

	if lineFormatter.HasPrefix("X") {
		lineFormatter.WriteString("X")
		lineFormatter.TrimPrefix("X")
	} else {
		lineFormatter.WriteString(" ")
	}
	lineFormatter.TrimSpace()

	if lineFormatter.HasPrefix("]") {
		lineFormatter.WriteString("]")
		lineFormatter.TrimPrefix("]")
	} else {
		return ""
	}
	lineFormatter.WriteString(" ")
	lineFormatter.TrimSpace()

	lineFormatter.WriteString(lineFormatter.Line)
	return lineFormatter.String()
}
