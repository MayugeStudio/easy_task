package code

import "strings"

func ParseTaskStringsToTasks(taskStrings []string) []*Task {
	tasks := make([]*Task, 0)
	taskStrings = FormatTaskStrings(taskStrings)
	taskLines := createLineHandlersFromStrings(taskStrings)
	for _, line := range taskLines {
		task := ParseTaskLineToTask(line)
		tasks = append(tasks, task)
	}
	return tasks
}

func ParseTaskLineToTask(line LineHandler) *Task {
	task := NewTask()
	tokens := line.Tokenize()
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
	return task
}

func createLineHandlersFromStrings(rawTaskStrings []string) []LineHandler {
	result := make([]LineHandler, 0)
	var lineHandler LineHandler
	for _, str := range rawTaskStrings {
		lineHandler = createLineHandlerFromString(str)
		result = append(result, lineHandler)
	}
	return result
}

func createLineHandlerFromString(rawString string) LineHandler {
	str := strings.TrimPrefix(rawString, "-")
	str = strings.TrimSpace(str)

	if strings.HasPrefix(str, "[") {
		return &taskLine{str: rawString}
	}

	return &groupLine{str: rawString}
}

type LineHandler interface {
	Tokenize() []string
	String() string
}

type taskLine struct {
	str string
}

func (t *taskLine) Tokenize() []string {
	panic("implement me")
}

func (t *taskLine) String() string {
	return t.str
}

type groupLine struct {
	str string
}

func (g *groupLine) Tokenize() []string {
	panic("implement me")
}

func (g *groupLine) String() string {
	return g.str
}
