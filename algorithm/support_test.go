package algorithm

import (
	"testing"

	"github.com/ready-steady/adapt/grid/equidistant"
	"github.com/ready-steady/adapt/internal"
	"github.com/ready-steady/assert"
)

func TestIsAdmissible(t *testing.T) {
	const (
		ni = 2
	)

	cases := []struct {
		levels []uint64
		orders []uint64
		result bool
	}{
		{
			[]uint64{
				0, 0,
				0, 1,
				1, 0,
				1, 1,
				1, 2,
			},
			[]uint64{
				0, 0,
				0, 2,
				2, 0,
				2, 2,
				2, 3,
			},
			true,
		},
		{
			[]uint64{
				0, 0,
				0, 1,
				1, 0,
				1, 1,
				1, 2,
			},
			[]uint64{
				0, 0,
				0, 2,
				2, 0,
				2, 2,
				2, 1,
			},
			false,
		},
		{
			[]uint64{
				0, 0,
				0, 1,
				1, 0,
				1, 1,
				2, 2,
			},
			[]uint64{
				0, 0,
				0, 2,
				2, 0,
				2, 2,
				3, 3,
			},
			false,
		},
	}

	for _, c := range cases {
		indices := internal.Compose(c.levels, c.orders)
		assert.Equal(IsAdmissible(indices, ni, equidistant.ClosedParent), c.result, t)
	}
}

func TestIsUnique(t *testing.T) {
	const (
		ni = 2
	)

	var levels, orders []uint64

	levels = []uint64{
		1, 2,
		3, 4,
		5, 6,
	}
	orders = []uint64{
		6, 5,
		4, 3,
		2, 1,
	}
	assert.Equal(IsUnique(internal.Compose(levels, orders), ni), true, t)

	levels = []uint64{
		1, 2,
		3, 4,
		5, 6,
		1, 2,
	}
	orders = []uint64{
		6, 5,
		4, 3,
		2, 1,
		6, 5,
	}
	assert.Equal(IsUnique(internal.Compose(levels, orders), ni), false, t)
}