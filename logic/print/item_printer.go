package print

import (
	"fmt"
	"io"
	"strings"

	"github.com/MayugeStudio/easy_task/domain"
)

const (
	doneSymbol   = "X"
	undoneSymbol = " "
)

func Items(w io.Writer, items domain.Items) error {
	for _, item := range items {
		str := getItemString(item, 0)
		if _, err := fmt.Fprint(w, str); err != nil {
			return fmt.Errorf("printing item: %w", err)
		}
	}
	return nil
}

func getItemString(item domain.Item, indentLevel int) string {
	if !item.IsParent() {
		var doneStr string
		if item.Progress() == 1 { // TODO: Implement float64 constant in domain package?
			doneStr = doneSymbol
		} else {
			doneStr = undoneSymbol
		}
		indentStr := strings.Repeat(" ", indentLevel)
		return fmt.Sprintf("%s[%s] %s\n", indentStr, doneStr, item.Title())
	}

	var b strings.Builder
	indentStr := strings.Repeat(" ", indentLevel)
	progress := item.Progress()
	taskProgressStr := getProgressString(progress, 20)
	titleString := fmt.Sprintf("%s%s %s\n", indentStr, item.Title(), taskProgressStr)
	b.WriteString(titleString)

	for _, child := range item.Children() {
		childStr := getItemString(child, indentLevel+2)
		b.WriteString(childStr)
	}

	return b.String()
}
