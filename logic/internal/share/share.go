package share

import (
	"unicode"

	"github.com/MayugeStudio/easy_task/utils"
)

func IsGroupTitle(s string) bool {
	l := utils.NewLine(s).TrimSpace()
	if !l.HasPrefix("-") {
		return false
	}
	l = l.TrimPrefix("-").TrimSpace()
	return !l.HasPrefix("[")
}

func IsTaskString(s string) bool {
	l := utils.NewLine(s).TrimSpace()

	if !l.HasPrefix("-") {
		return false
	}
	l = l.TrimPrefix("-").TrimSpace()

	if !l.HasPrefix("[") {
		return false
	}
	l = l.TrimPrefix("[").TrimSpace()

	if l.HasPrefix("X") || l.HasPrefix("x") {
		l = l.TrimPrefix("X").TrimSpace()
	}

	if !l.HasPrefix("]") {
		return false
	}

	return true
}

func IsItemModificationString(s string) bool {
	l := utils.NewLine(s)
	l = l.TrimSpace()
	if !l.HasPrefix("<-") {
		return false
	}
	l = l.TrimPrefix("<-").TrimSpace()
	if !l.HasPrefix("[") {
		return false
	}
	l = l.TrimPrefix("[").TrimSpace() // TODO: Implement 'Are' and 'Is' methods to Line struct
	return true
}

func GetIndentLevel(str string) int {
	level := 0
	runes := []rune(str)
	for _, r := range runes {
		if unicode.IsSpace(r) {
			level++
			continue
		}
		break
	}
	return level
}
