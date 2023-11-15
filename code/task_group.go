package code

type TaskGroup struct {
	Title      string
	ChildTasks []*Task
}

func NewTaskGroup(title string) *TaskGroup {
	return &TaskGroup{
		Title:      title,
		ChildTasks: make([]*Task, 0),
	}
}
