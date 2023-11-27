package domain

type Items struct {
	items []Item
}

func NewItems() *Items {
	return &Items{items: make([]Item, 0)}
}

func (c *Items) AddItem(i Item) {
	c.items = append(c.items, i)
}

func (c *Items) GetItems() []Item {
	return c.items
}

func (c *Items) Progress() float64 {
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
