package global

// Metric is an accuracy metric.
type Metric interface {
	// Done checks if the accuracy requirements have been satiated.
	Done(active Set) bool

	// Push takes into account new information about the target function.
	Push(values, surpluses []float64)

	// Score assigns a score to a dimensional location.
	Score(*Location) float64
}

// BasicMetric is a basic accuracy metric.
type BasicMetric struct {
	no       uint
	absolute float64
	relative float64

	errors []float64
	lower  []float64
	upper  []float64
}

// NewMetric creates a basic accuracy metric.
func NewMetric(no uint, absolute, relative float64) *BasicMetric {
	return &BasicMetric{
		no:       no,
		absolute: absolute,
		relative: relative,

		lower: repeatFloat64(infinity, no),
		upper: repeatFloat64(-infinity, no),
	}
}

func (self *BasicMetric) Done(active Set) bool {
	no, errors := self.no, self.errors
	δ := threshold(self.lower, self.upper, self.absolute, self.relative)
	for i := range active {
		for j := uint(0); j < no; j++ {
			if errors[i*no+j] > δ[j] {
				return false
			}
		}
	}
	return true
}

func (self *BasicMetric) Push(values, surpluses []float64) {
	no := self.no
	for i, point := range values {
		j := uint(i) % no
		if self.lower[j] > point {
			self.lower[j] = point
		}
		if self.upper[j] < point {
			self.upper[j] = point
		}
	}
	self.errors = append(self.errors, error(surpluses, no)...)
}

func (self *BasicMetric) Score(location *Location) float64 {
	no := self.no
	score := 0.0
	for _, value := range location.Surpluses {
		if value < 0.0 {
			value = -value
		}
		score += value
	}
	return score / float64(uint(len(location.Surpluses))/no)
}

func error(surpluses []float64, no uint) []float64 {
	error := repeatFloat64(-infinity, no)
	for i, value := range surpluses {
		j := uint(i) % no
		if value < 0.0 {
			value = -value
		}
		if value > error[j] {
			error[j] = value
		}
	}
	return error
}

func threshold(lower, upper []float64, absolute, relative float64) []float64 {
	no := uint(len(lower))
	threshold := make([]float64, no)
	for i := uint(0); i < no; i++ {
		threshold[i] = relative * (upper[i] - lower[i])
		if threshold[i] < absolute {
			threshold[i] = absolute
		}
	}
	return threshold
}
