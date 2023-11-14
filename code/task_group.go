package code

type TaskGroup struct {
	Title      string
	ChildTasks []TaskPtr
}

func NewTaskGroup(title string) *TaskGroup {
	return &TaskGroup{
		Title:      title,
		ChildTasks: make([]TaskPtr, 0),
	}
}
