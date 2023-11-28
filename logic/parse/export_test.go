package parse

import "github.com/MayugeStudio/easy_task/domain"

var newTask = domain.NewTask
var newGroup = func(title string, items []domain.Item) *domain.Group {
	g := domain.NewGroup(title)
	for _, item := range items {
		g.AddItem(item)
	}
	return g
}
