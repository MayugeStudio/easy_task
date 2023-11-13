package code

type Task struct {
	Title  string
	IsDone bool
}

type TaskPtr *Task

func NewTask() TaskPtr {
	return &Task{
		Title:  "",
		IsDone: false,
	}
}
