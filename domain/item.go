package domain

type Item interface {
	Title() string
	Label() string
	Priority() Priority
	Estimate() EstimateTime
	Progress() float64
	IsParent() bool
	Children() []Item
}
type EstimateTime int
type Priority int
