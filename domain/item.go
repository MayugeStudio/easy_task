package domain

type Priority int
type EstimateTime int

type Item interface {
	Title() string
	Label() string
	Priority() Priority
	Estimate() EstimateTime
	Progress() float64
	IsParent() bool
	Children() []Item
}

type Items []Item

func NewItems() Items {
	return make([]Item, 0)
}
func (i Items) Progress() float64 {
	if len(i) == 0 {
		return 0
	}

	var sum float64
	totalNum := len(i)
	for _, item := range i {
		sum += item.Progress()
	}

	return sum / float64(totalNum)
}
