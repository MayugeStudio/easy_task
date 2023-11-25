package domain

type Task struct {
	Title  string
	IsDone bool
}

func NewTask(title string, isDone bool) *Task {
	return &Task{
		Title:  title,
		IsDone: isDone,
	}
}

func (t *Task) Progress() float64 {
	if t.IsDone {
		return 1
	}
	return 0
}
