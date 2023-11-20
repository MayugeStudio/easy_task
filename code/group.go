package code

type Group struct {
	tasks []*Task
}

func NewGroup() *Group {
	return &Group{
		tasks: make([]*Task, 0),
	}
}

func (g *Group) AddTask(t *Task) {
	g.tasks = append(g.tasks, t)
}

func (g *Group) GetTasks() []*Task {
	return g.tasks
}
