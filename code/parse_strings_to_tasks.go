package code

import (
	"fmt"
	"strings"
)

func ParseStringsToTasks(lines []string) []*Task {
	lines = FormatTaskStrings(lines)
	tasks := make([]*Task, 0, len(lines))
	// Extruct Group Lines : 1
	// Parse Group Lines : 2
	// Parse Single Task Lines : 3
	for _, line := range lines {
		if strings.HasPrefix(line, "-") {
			// Trim unnecessary tokens.
			line = strings.TrimPrefix(line, "- ")
			line = strings.Replace(line, "[", "", 1)
			line = strings.Replace(line, "]", "", 1)
			tokens := strings.Fields(line)
			// Process each token until the tokens slice is empty.
			title := ""
			isDone := false
			for len(tokens) > 0 {
				token := tokens[0]
				switch token {
				case "X":
					isDone = true
				default:
					title = strings.Join(tokens, " ")
					tokens = nil
				}
				if len(tokens) > 0 {
					tokens = tokens[1:]
				}
			}
			task := NewTask(title, isDone)
			tasks = append(tasks, task)
		} else {
			continue
		}
	}
	noSpaceLine := strings.TrimSpace(line)
	if !strings.HasPrefix(noSpaceLine, "-") {
		fmt.Printf("!strings.HasPrefix(%q, \"-\") == true\n", line)
		return false
	}
	return true
}
