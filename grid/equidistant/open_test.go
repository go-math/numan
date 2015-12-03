package equidistant

import (
	"testing"

	"github.com/ready-steady/assert"
)

func TestOpenCompute1D(t *testing.T) {
	grid := NewOpen(1)

	levels := []uint64{0, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3}
	orders := []uint64{0, 0, 2, 0, 2, 4, 6, 0, 2, 4, 6, 8, 10, 12, 14}
	nodes := []float64{
		0.5000, 0.2500, 0.7500, 0.1250, 0.3750,
		0.6250, 0.8750, 0.0625, 0.1875, 0.3125,
		0.4375, 0.5625, 0.6875, 0.8125, 0.9375,
	}

	assert.Equal(grid.Compute(compose(levels, orders)), nodes, t)
}

func TestOpenCompute2D(t *testing.T) {
	grid := NewOpen(2)

	levels := []uint64{
		0, 0,

		0, 1,
		0, 1,
		0, 2,
		0, 2,
		0, 2,
		0, 2,

		1, 0,
		1, 0,

		1, 1,
		1, 1,
		1, 1,
		1, 1,

		1, 2,
		1, 2,
		1, 2,
		1, 2,
		1, 2,
		1, 2,
		1, 2,
		1, 2,

		2, 0,
		2, 0,
		2, 0,
		2, 0,

		2, 1,
		2, 1,
		2, 1,
		2, 1,
		2, 1,
		2, 1,
		2, 1,
		2, 1,

		2, 2,
		2, 2,
		2, 2,
		2, 2,
		2, 2,
		2, 2,
		2, 2,
		2, 2,
		2, 2,
		2, 2,
		2, 2,
		2, 2,
		2, 2,
		2, 2,
		2, 2,
		2, 2,
	}

	orders := []uint64{
		0, 0,

		0, 0,
		0, 2,
		0, 0,
		0, 2,
		0, 4,
		0, 6,

		0, 0,
		2, 0,

		0, 0,
		0, 2,
		2, 0,
		2, 2,

		0, 0,
		0, 2,
		0, 4,
		0, 6,
		2, 0,
		2, 2,
		2, 4,
		2, 6,

		0, 0,
		2, 0,
		4, 0,
		6, 0,

		0, 0,
		0, 2,
		2, 0,
		2, 2,
		4, 0,
		4, 2,
		6, 0,
		6, 2,

		0, 0,
		0, 2,
		0, 4,
		0, 6,
		2, 0,
		2, 2,
		2, 4,
		2, 6,
		4, 0,
		4, 2,
		4, 4,
		4, 6,
		6, 0,
		6, 2,
		6, 4,
		6, 6,
	}

	nodes := []float64{
		0.500, 0.500,

		0.500, 0.250,
		0.500, 0.750,
		0.500, 0.125,
		0.500, 0.375,
		0.500, 0.625,
		0.500, 0.875,

		0.250, 0.500,
		0.750, 0.500,

		0.250, 0.250,
		0.250, 0.750,
		0.750, 0.250,
		0.750, 0.750,

		0.250, 0.125,
		0.250, 0.375,
		0.250, 0.625,
		0.250, 0.875,
		0.750, 0.125,
		0.750, 0.375,
		0.750, 0.625,
		0.750, 0.875,

		0.125, 0.500,
		0.375, 0.500,
		0.625, 0.500,
		0.875, 0.500,

		0.125, 0.250,
		0.125, 0.750,
		0.375, 0.250,
		0.375, 0.750,
		0.625, 0.250,
		0.625, 0.750,
		0.875, 0.250,
		0.875, 0.750,

		0.125, 0.125,
		0.125, 0.375,
		0.125, 0.625,
		0.125, 0.875,
		0.375, 0.125,
		0.375, 0.375,
		0.375, 0.625,
		0.375, 0.875,
		0.625, 0.125,
		0.625, 0.375,
		0.625, 0.625,
		0.625, 0.875,
		0.875, 0.125,
		0.875, 0.375,
		0.875, 0.625,
		0.875, 0.875,
	}

	assert.Equal(grid.Compute(compose(levels, orders)), nodes, t)
}

func TestOpenChildren1D(t *testing.T) {
	grid := NewOpen(1)

	levels := []uint64{0, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3}
	orders := []uint64{0, 0, 2, 0, 2, 4, 6, 0, 2, 4, 6, 8, 10, 12, 14}
	childLevels := []uint64{
		1, 1,
		2, 2, 2, 2,
		3, 3, 3, 3, 3, 3, 3, 3,
		4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	}
	childOrders := []uint64{
		0, 2,
		0, 2, 4, 6,
		0, 2, 4, 6, 8, 10, 12, 14,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	}

	indices := grid.Children(compose(levels, orders))

	assert.Equal(indices, compose(childLevels, childOrders), t)
}

func TestOpenParent(t *testing.T) {
	grid := NewOpen(1)

	children := compose(
		[]uint64{0, 1, 1, 2, 2, 2, 2},
		[]uint64{0, 0, 2, 0, 2, 4, 6},
	)

	parents := compose(
		[]uint64{0, 0, 0, 1, 1, 1, 1},
		[]uint64{0, 0, 0, 0, 0, 2, 2},
	)

	for i := range children {
		grid.Parent(children[i:i+1], 0)
	}

	assert.Equal(children, parents, t)
}

func TestOpenSibling(t *testing.T) {
	grid := NewOpen(1)

	indices := compose(
		[]uint64{0, 1, 1, 2, 2, 2, 2},
		[]uint64{0, 0, 2, 0, 2, 4, 6},
	)

	siblings := compose(
		[]uint64{0, 1, 1, 2, 2, 2, 2},
		[]uint64{0, 2, 0, 2, 0, 6, 4},
	)

	for i := range indices {
		grid.Sibling(indices[i:i+1], 0)
	}

	assert.Equal(indices, siblings, t)
}