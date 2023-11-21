package src

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

type GroupContainer struct {
	groups []*Group
}

func NewGroupContainer() *GroupContainer {
	return &GroupContainer{
		groups: make([]*Group, 0),
	}
}

func (c *GroupContainer) AddGroup(g *Group) {
	c.groups = append(c.groups, g)
}
