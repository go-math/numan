package global

import (
	"math"

	"github.com/ready-steady/adapt/algorithm/external"
	"github.com/ready-steady/adapt/algorithm/internal"
)

// Strategy guides the interpolation process.
type Strategy interface {
	// Continue decides if the interpolation process should go on.
	Continue(*external.Active) bool

	// Push takes into account a new interpolation element and its score.
	Push(*Element, float64)

	// Select chooses an active index for refinement.
	Select(*external.Active) uint
}

type defaultStrategy struct {
	ni uint
	no uint

	εa float64
	εr float64

	scores []float64
	errors []float64

	lower []float64
	upper []float64
}

func newStrategy(ni, no uint, absolute, relative float64) *defaultStrategy {
	return &defaultStrategy{
		ni: ni,
		no: no,

		εa: absolute,
		εr: relative,

		lower: internal.RepeatFloat64(math.Inf(1.0), no),
		upper: internal.RepeatFloat64(math.Inf(-1.0), no),
	}
}

func (self *defaultStrategy) Continue(active *external.Active) bool {
	no, errors := self.no, self.errors
	ne := uint(len(errors)) / no
	if ne == 0 {
		return true
	}
	δ := threshold(self.lower, self.upper, self.εa, self.εr)
	for i := range active.Positions {
		if i >= ne {
			continue
		}
		for j := uint(0); j < no; j++ {
			if errors[i*no+j] > δ[j] {
				return true
			}
		}
	}
	return false
}

func (self *defaultStrategy) Push(element *Element, score float64) {
	self.updateBounds(element.Values)
	self.scores = append(self.scores, score)
	self.errors = append(self.errors, error(element.Surpluses, self.no)...)
}

func (self *defaultStrategy) Select(active *external.Active) uint {
	return internal.LocateMaxFloat64s(self.scores, active.Positions)
}

func (self *defaultStrategy) updateBounds(values []float64) {
	no := self.no
	for i, point := range values {
		j := uint(i) % no
		self.lower[j] = math.Min(self.lower[j], point)
		self.upper[j] = math.Max(self.upper[j], point)
	}
}

func error(surpluses []float64, no uint) []float64 {
	error := make([]float64, no)
	for i, value := range surpluses {
		j := uint(i) % no
		error[j] = math.Max(error[j], math.Abs(value))
	}
	return error
}

func threshold(lower, upper []float64, εa, εr float64) []float64 {
	threshold := make([]float64, len(lower))
	for i := range threshold {
		threshold[i] = math.Max(εr*(upper[i]-lower[i]), εa)
	}
	return threshold
}
