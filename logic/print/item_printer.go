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

func Items(w io.Writer, items []domain.Item) error {
	for _, item := range items {
		str := getItemString(item)
		if _, err := fmt.Fprint(w, str); err != nil {
			return fmt.Errorf("printing item: %w", err)
		}
	}
	return nil
}

func getItemString(item domain.Item) string {
	if !item.IsParent() {
		var doneStr string
		if item.Progress() == 1 { // TODO: Implement float64 constant in domain package?
			doneStr = doneSymbol
		} else {
			doneStr = undoneSymbol
		}
		return fmt.Sprintf("[%s] %s\n", doneStr, item.Title())
	}

	var b strings.Builder
	titleString := fmt.Sprintf("%s\n", item.Title())
	b.WriteString(titleString)

	for _, child := range item.Children() {
		childStr := fmt.Sprintf("  %s", getItemString(child))
		b.WriteString(childStr)
	}

	progress := item.Progress()
	taskProgressString := fmt.Sprintf("  %s\n", getProgressString(progress, 20))
	b.WriteString(taskProgressString)

	return b.String()
}
