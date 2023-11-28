package format

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/MayugeStudio/easy_task/utils"
)

func toFormattedModificationString(s string) (string, error) {
	indentCount := 0
	rs := []rune(s)
	for _, r := range rs {
		if unicode.IsSpace(r) {
			indentCount++
			continue
		}
		break
	}
	l := utils.NewLine(s).TrimSpace()
	if !l.HasPrefix("<-") {
		return "", errInvalidModifier
	}
	l = l.TrimPrefix("<-").TrimSpace()

	if !l.HasPrefix("[") {
		return "", fmt.Errorf("%w: while formatting modification string", errNoBracketStart)
	}
	l = l.TrimPrefix("[").TrimSpace()

	var attribute string
	if !l.HasPrefix("Tag") { // TODO: This implementation is temporary.
		return "", fmt.Errorf("%w: while formatting modification string", errInvalidModifierAttribute)
	}
	attribute = "Tag"
	l = l.TrimPrefix("Tag").TrimSpace() // TODO: Implement TrimPrefixInSlice that trim the prefix in passed slice to Line type.
	if !l.HasPrefix("]") {
		return "", fmt.Errorf("%w: while formatting modification string", errNoBracketEnd)
	}
	l = l.TrimPrefix("]").TrimSpace()

	if !l.HasPrefix(":") {
		return "", fmt.Errorf("%w: while formatting modification string", errNoColon)
	}
	l = l.TrimPrefix(":").TrimSpace()

	attributeValue := string(l)

	return fmt.Sprintf("%s<- [%s]: %s", strings.Repeat(" ", indentCount), attribute, attributeValue), nil
}
