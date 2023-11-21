package code

type TodoItemContainer struct {
	taskContainer  *TaskContainer
	groupContainer *GroupContainer
}

func NewTodoItemContainer() *TodoItemContainer {
	return &TodoItemContainer{
		taskContainer:  NewTaskContainer(),
		groupContainer: NewGroupContainer(),
	}
}

func (c *TodoItemContainer) AddTask(t *Task) {
	c.taskContainer.AddTask(t)
}

func (c *TodoItemContainer) AddGroup(g *Group) {
	c.groupContainer.AddGroup(g)
}

func (c *TodoItemContainer) GetTasks() []*Task {
	return c.taskContainer.tasks
}

func (c *TodoItemContainer) GetGroups() []*Group {
	return c.groupContainer.groups
}
