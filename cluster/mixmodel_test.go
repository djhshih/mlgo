package cluster

import (
	"testing"
)

var mixModelTests = []struct {
	x Matrix
	k int
	index Partitions
	means Matrix
}{
	{
		Matrix{
			{-10, -10},
			{-10,  -8},
			{ -8,  -8},
			{ -8, -10},
			{ 10,  10},
			{ 10,   8},
			{  8,   8},
			{  8,  10},
		},
		2,
		Partitions{0, 0, 0, 0, 1, 1, 1, 1},
		Matrix{
			{-9, -9},
			{9, 9},
		},
	},
}

func TestMixModel(t *testing.T) {
	for i, test := range mixModelTests {
		c := MixModel{X:test.x}
		classes := c.Cluster(test.k)
		if !classes.Index.Equal(test.index) {
			t.Errorf("#%d MixModel.Cluster(%d) got %v, want %v", i, test.k, classes.Index, test.index)
		}
		if !CoordinatesSetEqual(c.Means, test.means) {
			t.Errorf("#%d MixModel.Cluster(%d) got %v, want %v", i, test.k, c.Means, test.means)
		}
	}
}

