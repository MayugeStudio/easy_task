package code

import "strings"

type LineFormatter struct {
	Line string
}

func NewLineFormatter(line string) *LineFormatter {
	return &LineFormatter{Line: line}
}

func (f *LineFormatter) HasPrefix(prefix string) bool {
	return strings.HasPrefix(f.Line, prefix)
}

func (f *LineFormatter) TrimPrefix(prefix string) *LineFormatter {
	upperPrefix := strings.ToUpper(prefix)
	lowerPrefix := strings.ToLower(prefix)

	if strings.HasPrefix(f.Line, upperPrefix) {
		f.Line = strings.TrimPrefix(f.Line, upperPrefix)
	} else if strings.HasPrefix(f.Line, lowerPrefix) {
		f.Line = strings.TrimPrefix(f.Line, lowerPrefix)
	}

	return f
}

func (f *LineFormatter) Replace(old string, new string, n int) *LineFormatter {
	f.Line = strings.Replace(f.Line, old, new, n)
	return f
}

func (f *LineFormatter) TrimSpace() {
	f.Line = strings.TrimSpace(f.Line)
}
