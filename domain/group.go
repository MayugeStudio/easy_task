package domain

type Group struct {
	Title string
	Tasks []*Task
}

func NewGroup(title string) *Group {
	return &Group{
		Title: title,
		Tasks: make([]*Task, 0),
	}
}

func (g *Group) AddTask(t *Task) {
	g.Tasks = append(g.Tasks, t)
}

func (g *Group) Progress() float64 {
	if len(g.Tasks) == 0 {
		return 0
	}
	var sum float64
	for _, task := range g.Tasks {
		sum += task.Progress()
	}
	return sum / float64(len(g.Tasks))
}
