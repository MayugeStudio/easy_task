package code

type GroupContainer struct {
	groups []*Group
}

func NewGroupContainer() *GroupContainer {
	return &GroupContainer{
		groups: make([]*Group, 0),
	}
}

func (c *GroupContainer) AddGroup(g *Group) {
	c.groups = append(c.groups, g)
}
