package cluster

import (
	"testing"
)

var hClustersGenericTests = []struct{
	x Matrix
	metric MetricOp
	method int
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
		single_linkage,
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
		complete_linkage,
		2,
		Partitions{1, 1, 2, 2, 1, 1, 2, 2},
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
		average_linkage,
		4,
		Partitions{1, 1, 2, 2, 3, 3, 4, 4},
	},
}

func TestHClustersGeneric(t *testing.T) {
	for i, test := range hClustersGenericTests {
		c := NewHClustersGeneric(test.x, test.metric, test.method)
		classes := c.Cluster(test.k)
		if !classes.Index.Equal(test.index) {
			t.Errorf("#%d HClustersGeneric.Cluster(%d) got %v, want %v", i, test.k, classes.Index, test.index)
		}
	}
}

