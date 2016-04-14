package external

// Element contains information about an interpolation element.
type Element struct {
	Index []uint64 // Nodal index

	Volume  float64   // Basis-function volume
	Value   []float64 // Target-function value
	Surplus []float64 // Hierarchical surplus
}

// State contains information about an interpolation iteration.
type State struct {
	Lindices []uint64 // Level indices
	Indices  []uint64 // Nodal indices
	Counts   []uint   // Number of nodal indices for each level index

	Nodes     []float64 // Grid nodes
	Volumes   []float64 // Basis-function volumes
	Values    []float64 // Target-function values
	Estimates []float64 // Approximated values
	Surpluses []float64 // Hierarchical surpluses
	Scores    []float64 // Nodal-index scores
}

// Strategy controls the interpolation process.
type Strategy interface {
	// First returns the initial state of the first iteration.
	First() *State

	// Check returns true if the interpolation process should continue.
	Check(*State, *Surrogate) bool

	// Score assigns a score to an interpolation element.
	Score(*Element) float64

	// Next consumes the result of the current iteration and returns the initial
	// state of the next one.
	Next(*State, *Surrogate) *State
}
