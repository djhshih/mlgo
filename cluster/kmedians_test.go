package cluster

import (
	"testing"
)

var medianTests = []struct {
	in Vector
	ans float64
}{
	{Vector{1, 2, 3, 4, 5}, 3},
	{Vector{3, 2, 1, 5, 4}, 3},
	{Vector{1, 1, 4, 5}, 2.5},
}

func TestMedian(t *testing.T) {
	for i, test := range medianTests {
		out := median(test.in)
		if out != test.ans {
			t.Errorf("#%d median(%v) got %v, want %v", i, test.in, out, test.ans)
		}
	}
}

var kmediansTests = []struct {
	x Matrix
	metric MetricOp
	k int
	index Partitions
	centers Matrix
}{
	{
		Matrix{
			{-20, -30},
			{-10, -20},
			{-10, -18},
			{ -8, -18},
			{ -8, -20},
			{ 10,  20},
			{ 10,  18},
			{  8,  18},
			{  8,  20},
		},
		Euclidean,
		2,
		Partitions{0, 0, 0, 0, 0, 1, 1, 1, 1},
		Matrix{
			{-10, -20},
			{9, 19},
		},
	},
}

func TestKMedians(t *testing.T) {
	for i, test := range kmediansTests {
		c := NewKMedians(test.x, test.metric)
		classes := c.Cluster(test.k)
		if !classes.Index.Equal(test.index) {
			t.Errorf("#%d KMedians.Cluster(...) got %v, want %v", i, classes.Index, test.index)
		}
		if !CoordinatesSetEqual(c.Centers, test.centers) {
			t.Errorf("#%d KMedians.Cluster(...) got %v, want %v", i, c.Centers, test.centers)
		}
	}
}

