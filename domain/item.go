package domain

type Item interface {
	Title() string
	Label() string
	Priority() Priority
	Estimate() EstimateTime
	Progress() float64
}
type EstimateTime int
type Priority int

type ProgressItem interface {
	Progress() float64
}
