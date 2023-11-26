package domain

type Task struct {
	title  string
	isDone bool
}

func NewTask(title string, isDone bool) *Task {
	return &Task{
		title:  title,
		isDone: isDone,
	}
}

func (t *Task) Title() string {
	return t.title
}

func (t *Task) Label() string {
	//TODO implement me
	panic("implement me")
}

func (t *Task) Priority() Priority {
	//TODO implement me
	panic("implement me")
}

func (t *Task) Estimate() EstimateTime {
	//TODO implement me
	panic("implement me")
}

func (t *Task) Progress() float64 {
	if t.isDone {
		return 1
	}
	return 0
}

func (t *Task) IsParent() bool {
	return false
}

func (t *Task) Children() []Item {
	return nil
}
