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
