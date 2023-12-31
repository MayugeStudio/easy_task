package print

import (
	"fmt"
	"io"
	"strings"

	"github.com/MayugeStudio/easy_task/domain"
)

const (
	progressSymbol           = "#"
	defaultProgressBarLength = 40.0
)

func Progress(w io.Writer, items domain.Items) error {
	progress := items.Progress()
	progressString := getProgressString(progress, defaultProgressBarLength)
	if _, err := fmt.Fprint(w, progressString); err != nil {
		return err
	}
	return nil
}

func getProgressString(progress, length float64) string {
	barLength := int(progress * length)
	doneStr := strings.Repeat(progressSymbol, barLength)
	undoneStr := strings.Repeat(" ", int(length)-barLength)
	return fmt.Sprintf("[%s%s]%.1f%%", doneStr, undoneStr, progress*100)
}
