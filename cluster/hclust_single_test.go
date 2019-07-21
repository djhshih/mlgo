package cluster

import (
	"testing"
)

var hClustersSingleTests = []struct{
	x Matrix
	metric MetricOp
	k int
	index Partitions
}{
	{
		Matrix{
			{0, 0},
			{1, 0},
			{2, 0},
			{3, 0},
			{0, 3},
			{1, 3},
			{2, 3},
			{3, 3},
		},
		Euclidean,
		2,
		Partitions{1, 1, 1, 1, 2, 2, 2, 2},
	},
	{
		Matrix{
			{0, 0},
			{1, 0},
			{3, 0},
			{4, 0},
			{0, 4},
			{1, 4},
			{3, 4},
			{4, 4},
		},
		Euclidean,
		2,
		Partitions{1, 1, 1, 1, 2, 2, 2, 2},
	},
	{
		Matrix{
			{0, 0},
			{1, 0},
			{3, 0},
			{4, 0},
			{0, 4},
			{1, 4},
			{3, 4},
			{4, 4},
		},
		Euclidean,
		4,
		Partitions{1, 1, 2, 2, 3, 3, 4, 4},
	},
}

func TestHClustersSingle(t *testing.T) {
	for i, test := range hClustersSingleTests {
		c := NewHClustersSingle(test.x, test.metric, nil)
		classes := c.Cluster(test.k)
		if !classes.Index.Equal(test.index) {
			t.Errorf("#%d HClustersSingle.Cluster(%d) got %v, want %v", i, test.k, classes.Index, test.index)
		}
	}
}

