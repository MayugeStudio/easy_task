package print

import (
	"fmt"
	"io"
	"strings"

	"github.com/MayugeStudio/easy_task/domain"
)

func Groups(w io.Writer, groups []*domain.Group) error {
	for _, group := range groups {
		groupString := getGroupString(group)
		if _, err := fmt.Fprint(w, groupString); err != nil {
			return err
		}
	}
	return nil
}

func getGroupString(group *domain.Group) string {
	var b strings.Builder
	titleString := fmt.Sprintf("%s\n", group.Title())
	b.WriteString(titleString)

	length := getMaxTaskTitleLength(group.Tasks())
	for _, task := range group.Tasks() {
		taskString := fmt.Sprintf("  %s\n", getTaskString(task, length))
		b.WriteString(taskString)
	}
	progress := group.Progress()
	taskProgressString := fmt.Sprintf("  %s\n", getProgressString(progress, 20))
	b.WriteString(taskProgressString)
	return b.String()
}
