package code

type Task struct {
	Title  string
	IsDone bool
}

func NewTask() *Task {
	return &Task{
		Title:  "",
		IsDone: false,
	}
}
