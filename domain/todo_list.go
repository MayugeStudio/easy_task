package domain

type TodoList struct {
	items []Item
}

func NewTodoList() *TodoList {
	return &TodoList{items: make([]Item, 0)}
}

func (c *TodoList) AddItem(i Item) {
	c.items = append(c.items, i)
}

func (c *TodoList) GetItems() []Item {
	return c.items
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
