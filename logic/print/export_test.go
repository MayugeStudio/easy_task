package print

import "github.com/MayugeStudio/easy_task/domain"

var newTask = domain.NewTask

var newGroup = func(title string, tasks []*domain.Task) *domain.Group {
	g := domain.NewGroup(title)
	for _, task := range tasks {
		g.AddItem(task)
	}
	return g
}
