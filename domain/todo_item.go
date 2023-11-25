package domain

type TodoList struct {
	tasks         []*Task
	groups        []*Group
	doneTaskCount int
}

func NewTodoList() *TodoList {
	return &TodoList{
		tasks:         make([]*Task, 0),
		groups:        make([]*Group, 0),
		doneTaskCount: 0,
	}
}

func (c *TodoList) AddTask(t *Task) {
	c.tasks = append(c.tasks, t)
	if t.IsDone {
		c.doneTaskCount++
	}
}

func (c *TodoList) AddGroup(g *Group) {
	c.groups = append(c.groups, g)
}

func (c *TodoList) GetTasks() []*Task {
	return c.tasks
}

func (c *TodoList) GetGroups() []*Group {
	return c.groups
}
