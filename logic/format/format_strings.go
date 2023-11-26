package format

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/MayugeStudio/easy_task/logic/internal/share"
	"github.com/MayugeStudio/easy_task/utils"
)

// FIXME: Make the error message more detailed.
var (
	errSyntax                       = errors.New("format error")
	errNoDash                       = fmt.Errorf("%w: no dash", errSyntax)
	errNoBracketStart               = fmt.Errorf("%w: no bracket start", errSyntax)
	errNoBracketEnd                 = fmt.Errorf("%w: no bracket end", errSyntax)
	errNoColon                      = fmt.Errorf("%w: no colon", errSyntax)
	errInvalidIndent                = fmt.Errorf("%w: no valid indent", errSyntax) // FIXME: fix error message.
	errInvalidModification          = fmt.Errorf("%w: invalid modification", errSyntax)
	errInvalidModificationAttribute = fmt.Errorf("%w: invalid modification attribute", errSyntax)
)

func ToValidStrings(taskStrings []string) ([]string, []error) {
	result := make([]string, 0)
	errs := make([]error, 0)
	inGroup := false
	for _, str := range taskStrings {
		var formattedString string
		var err error
		if share.IsGroupTitle(str) { // TODO: rename share -> share/XXX. 'share.IsGroupTitle' is difficult to read.
			formattedString, err = toFormattedGroupTitle(str)
			inGroup = true
		} else if inGroup && share.IsGroupTaskString(str) {
			formattedString, err = toFormattedGroupTaskString(str)
		} else if share.IsSingleTaskString(str) {
			formattedString, err = toFormattedTaskString(str)
			inGroup = false
		} else if share.IsItemModificationString(str) {
			formattedString, err = toFormattedModificationString(str)
		} else {
			if !strings.HasPrefix(str, "  ") {
				inGroup = false
			}
			continue
		}
		if err != nil {
			errs = append(errs, err)
			continue
		}
		result = append(result, formattedString)
	}
	return result, errs
}

func toFormattedGroupTitle(s string) (string, error) {
	title, err := extractGroupTitle(s)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("- %s", title), nil
}

func toFormattedGroupTaskString(s string) (string, error) {
	if !strings.HasPrefix(s, " ") {
		return "", errInvalidIndent
	}
	noSpaceStr := strings.TrimSpace(s)
	formattedString, err := toFormattedTaskString(noSpaceStr)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("  %s", formattedString), nil
}

func toFormattedTaskString(s string) (string, error) {
	if !strings.HasPrefix(s, "-") {
		return "", errNoDash
	}

	l := utils.NewLine(s)

	l = l.TrimPrefix("-").TrimSpace()

	if !l.HasPrefix("[") {
		return "", errNoBracketStart
	}
	l = l.TrimPrefix("[").TrimSpace()

	statusStr, err := toFormattedTaskStatus(s)
	if err != nil {
		return "", err
	}
	l = l.TrimPrefix(statusStr).TrimSpace()

	if !l.HasPrefix("]") {
		return "", errNoBracketEnd
	}

	l = l.TrimPrefix("]").TrimSpace()

	return fmt.Sprintf("- [%s] %s", statusStr, l), nil
}

func toFormattedTaskStatus(s string) (string, error) {
	if !strings.HasPrefix(s, "-") {
		return "", errNoDash
	}

	l := utils.NewLine(s)

	l = l.TrimPrefix("-").TrimSpace()

	if !l.HasPrefix("[") {
		return "", errNoBracketStart
	}
	l = l.TrimPrefix("[").TrimSpace()

	if l.HasPrefix("X") || l.HasPrefix("x") {
		return "X", nil
	}
	// FIXME: Other strings are ignored.
	return " ", nil
}

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
		return "", errInvalidModification
	}
	l = l.TrimPrefix("<-").TrimSpace()

	if !l.HasPrefix("[") {
		return "", fmt.Errorf("%w: while formatting modification string", errNoBracketStart)
	}
	l = l.TrimPrefix("[").TrimSpace()

	var attribute string
	if !l.HasPrefix("Tag") { // TODO: This implementation is temporary.
		return "", fmt.Errorf("%w: while formatting modification string", errInvalidModificationAttribute)
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

func extractGroupTitle(s string) (string, error) {
	if !strings.HasPrefix(s, "-") {
		// FIXME: Define error var.
		return "", fmt.Errorf("%w: invalid group title %q", errSyntax, s)
	}
	s = strings.TrimPrefix(s, "-")
	return strings.TrimSpace(s), nil
}
