package format

import (
	"fmt"
	"strings"

	"github.com/MayugeStudio/easy_task/logic/internal/share"
	"github.com/MayugeStudio/easy_task/utils"
)

func toFormattedGroupTitle(s string) (string, error) {
	indentLevel := share.GetIndentLevel(s)
	indentStr := strings.Repeat(" ", indentLevel)
	title, err := extractGroupTitle(s)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s- %s", indentStr, title), nil
}

func toFormattedGroupTaskString(s string) (string, error) {
	l := utils.NewLine(s)
	if !l.HasPrefix(" ") {
		return "", errInvalidIndent
	}
	indentLevel := share.GetIndentLevel(l.String())
	// TODO: Change Indent rule. Use indentLevel % 2 == 1 -> indentLevel += 1
	if indentLevel == 1 {
		indentLevel = 2
	}
	indentStr := strings.Repeat(" ", indentLevel)
	fStr, err := toFormattedTaskString(l.TrimSpace().String())
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", indentStr, fStr), nil
}

func extractGroupTitle(s string) (string, error) {
	s = strings.TrimSpace(s)
	if !strings.HasPrefix(s, "-") {
		// FIXME: Define error var.
		return "", fmt.Errorf("%w: invalid group title %q", errSyntax, s)
	}
	s = strings.TrimPrefix(s, "-")
	return strings.TrimSpace(s), nil
}
