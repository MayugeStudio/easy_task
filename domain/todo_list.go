package domain

type TodoList struct {
	tasks         []*Task
	groups        []*Group
	items         []ProgressItem
	doneTaskCount int
}

func NewTodoList() *TodoList {
	return &TodoList{
		tasks:         make([]*Task, 0),
		groups:        make([]*Group, 0),
		items:         make([]ProgressItem, 0),
		doneTaskCount: 0,
	}
}

func (c *TodoList) AddTask(t *Task) {
	c.tasks = append(c.tasks, t)
	if t.IsDone {
		c.doneTaskCount++
	}
	c.items = append(c.items, t)
}

func (c *TodoList) AddGroup(g *Group) {
	c.groups = append(c.groups, g)
	c.items = append(c.items, g)
}

func (c *TodoList) GetTasks() []*Task {
	return c.tasks
}

func (c *TodoList) GetGroups() []*Group {
	return c.groups
}

func (c *TodoList) Progress() float64 {
	if len(c.items) == 0 {
		return 0
	}

	var sum float64
	totalNum := len(c.items)
	for _, item := range c.items {
		sum += item.Progress()
	}

	return sum / float64(totalNum)
}
