package domain

type Group struct {
	title string
	tasks []*Task
}

func NewGroup(title string) *Group {
	return &Group{
		title: title,
		tasks: make([]*Task, 0),
	}
}

func (g *Group) AddTask(t *Task) {
	g.tasks = append(g.tasks, t)
}

func (g *Group) Tasks() []*Task {
	return g.tasks
}

func (g *Group) Title() string {
	return g.title
}

func (g *Group) Label() string {
	//TODO implement me
	panic("implement me")
}

func (g *Group) Priority() Priority {
	//TODO implement me
	panic("implement me")
}

func (g *Group) Estimate() EstimateTime {
	//TODO implement me
	panic("implement me")
}

func (g *Group) Progress() float64 {
	if len(g.tasks) == 0 {
		return 0
	}
	var sum float64
	for _, task := range g.tasks {
		sum += task.Progress()
	}
	return sum / float64(len(g.tasks))
}
