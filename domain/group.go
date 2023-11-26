package domain

type Group struct {
	title string
	items []Item
}

func NewGroup(title string) *Group {
	return &Group{
		title: title,
		items: make([]Item, 0),
	}
}

func (g *Group) AddItem(i Item) {
	g.items = append(g.items, i)
}

func (g *Group) Title() string {
	return g.title
}

func (g *Group) Label() string {
	//TODO implement me
	panic("implement me")
}

func (g *Group) Priority() Priority {
	//TODO implement me
	panic("implement me")
}

func (g *Group) Estimate() EstimateTime {
	//TODO implement me
	panic("implement me")
}

func (g *Group) Progress() float64 {
	if len(g.items) == 0 {
		return 0
	}
	var sum float64
	for _, task := range g.items {
		sum += task.Progress()
	}
	return sum / float64(len(g.items))
}

func (g *Group) IsParent() bool {
	return true
}

func (g *Group) Children() []Item {
	return g.items
}
