package src

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

type TaskContainer struct {
	tasks         []*Task
	doneTaskCount int
}

func NewTaskContainer() *TaskContainer {
	return &TaskContainer{
		tasks:         make([]*Task, 0),
		doneTaskCount: 0,
	}
}

func (c *TaskContainer) AddTask(t *Task) {
	c.tasks = append(c.tasks, t)
	if t.IsDone {
		c.doneTaskCount++
	}
}
