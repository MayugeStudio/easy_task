package code

import (
	"fmt"
	"strings"
)

func isNestedTaskLine(line string) bool {
	if !strings.HasPrefix(line, " ") {
		fmt.Printf("!strings.HasPrefix(%q, \" \") == true\n", line)
		return false
	}
	noSpaceLine := strings.TrimSpace(line)
	if !strings.HasPrefix(noSpaceLine, "-") {
		fmt.Printf("!strings.HasPrefix(%q, \"-\") == true\n", line)
		return false
	}
	return true
}

func ParseStringsToTasks(lines []string) []*Task {
	lines = FormatTaskStrings(lines)
	tasks := make([]*Task, 0, len(lines))
	// Extruct TaskGroup Lines
	var groupLinesContainer [][]string
	var currentGroupLines []string
	var isGroup = false
	for i := 0; i < len(lines); i++ {
		if IsGroupTitle(lines[i]) && !isGroup {
			println("IsGroupTitle() = true")
			isGroup = true
			currentGroupLines = make([]string, 0)
		}

		if !isGroup {
			continue
		}
		currentGroupLines = append(currentGroupLines, lines[i])

		// Check for the last line or the next line is not nested.
		isLastLine := i == len(lines)-1
		isNextLineNotNested := i <= len(lines)-1 && !isNestedTaskLine(lines[i+1])
		if isLastLine || isNextLineNotNested {
			fmt.Println("isLastLine || isNextLineNotNested == true")
			fmt.Println("isLastLine =", isLastLine)
			fmt.Println("isNextLineNotNested =", isNextLineNotNested)
			groupLinesContainer = append(groupLinesContainer, currentGroupLines)
			isGroup = false
		}
	}
	for i, l := range groupLinesContainer {
		fmt.Printf(">>%d<<\n", i)
		fmt.Printf("  [%s]\n", strings.Join(l, ", "))
	}
	fmt.Println("--------------------")

	// Parse GroupLines Lines
	// Parse Single Task Lines
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
	return tasks
}
