package model

type Report struct {
	Name         string
	Code         string
	Rate         float64
	Share        float64
	Contribution float64
	Children     []*Report
}
