package domain

type TodoItemContainer struct {
	tasks         []*Task
	groups        []*Group
	doneTaskCount int
}

func NewTodoItemContainer() *TodoItemContainer {
	return &TodoItemContainer{
		tasks:         make([]*Task, 0),
		groups:        make([]*Group, 0),
		doneTaskCount: 0,
	}
}

func (c *TodoItemContainer) AddTask(t *Task) {
	c.tasks = append(c.tasks, t)
	if t.IsDone {
		c.doneTaskCount++
	}
}

func (c *TodoItemContainer) AddGroup(g *Group) {
	c.groups = append(c.groups, g)
}

func (c *TodoItemContainer) GetTasks() []*Task {
	return c.tasks
}

func (c *TodoItemContainer) GetGroups() []*Group {
	return c.groups
}
