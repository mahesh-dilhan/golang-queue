package queue

type CountrySeverity int

const (
	Mild CountrySeverity = iota
	Moderate
	Critial
)

type Country struct {
	name     string
	severity CountrySeverity
}
