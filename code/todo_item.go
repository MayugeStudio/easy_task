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
