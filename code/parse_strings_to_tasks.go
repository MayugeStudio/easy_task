package code

import "strings"

func ParseStringsToTasks(lines []string) []*Task {
	tasks := make([]*Task, 0)
	lines = FormatTaskStrings(lines)
	for _, line := range lines {
		if strings.HasPrefix(line, "-") {
			line = strings.TrimPrefix(line, "-")
			line = strings.TrimSpace(line)
			line = strings.Replace(line, "[", "", 1)
			line = strings.Replace(line, "]", "", 1)
			line = strings.TrimSpace(line)
			tokens := strings.Fields(line)
			task := NewTask()
			// Process each token until the tokens slice is empty.
			for len(tokens) > 0 {
				token := tokens[0]
				switch token {
				case "X":
					task.IsDone = true
				default:
					task.Title = strings.Join(tokens, " ")
					tokens = nil
				}
				if len(tokens) > 0 {
					tokens = tokens[1:]
				}
			}
			tasks = append(tasks, task)
		} else {
			continue
		}
	}
	return tasks
}
