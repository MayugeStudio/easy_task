package format

import (
	"fmt"
	"strings"

	"github.com/MayugeStudio/easy_task/logic/internal/share"
)

func toFormattedGroupTitle(s string) (string, error) {
	indentLevel := share.GetIndentLevel(s)
	title, err := extractGroupTitle(s)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s- %s", strings.Repeat(" ", indentLevel), title), nil
}

func toFormattedGroupTaskString(s string) (string, error) {
	if !strings.HasPrefix(s, " ") {
		return "", errInvalidIndent
	}
	indentLevel := share.GetIndentLevel(s)
	noSpaceStr := strings.TrimSpace(s)
	formattedString, err := toFormattedTaskString(noSpaceStr)
	if err != nil {
		return "", err
	}
	if indentLevel == 1 { // TODO: Change Indent rule. Use indentLevel % 2 == 1 -> indentLevel += 1
		indentLevel = 2
	}
	return fmt.Sprintf("%s%s", strings.Repeat(" ", indentLevel), formattedString), nil
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
