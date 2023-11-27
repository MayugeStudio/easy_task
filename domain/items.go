package domain

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
