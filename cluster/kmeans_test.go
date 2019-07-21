package cluster

import (
	"testing"
	"math/rand"
)

var kmeansTests = []struct {
	x Matrix
	metric MetricOp
	k int
	partitions Partitions
	centers Matrix
}{
	{
		Matrix{
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
		Partitions{0, 0, 0, 0, 1, 1, 1, 1},
		Matrix{
			{-9, -19},
			{9, 19},
		},
	},
}

func TestKMeans(t *testing.T) {
	for i, test := range kmeansTests {
		c := NewKMeans(test.x, test.metric)
		classes := c.Cluster(test.k)
		if !classes.Index.Equal(test.partitions) {
			t.Errorf("#%d KMeans.Cluster(...) got %v, want %v", i, classes.Index, test.partitions)
		}
		if !CoordinatesSetEqual(c.Centers, test.centers) {
			t.Errorf("#%d KMeans.Cluster(...) got %v, want %v", i, c.Centers, test.centers)
		}

		// test if same results are obtained after permutation

		index := rand.Perm(c.Len())
		d := c.Subset(index)
		classes2 := d.Cluster(test.k)
		// inverse the permutation
		InvPerm(index)
		// get the corrected index of the resulting classes
		partitions := Partitions( Permute(classes2.Index, index) )
		if !partitions.Equal(test.partitions) {
			t.Errorf("#%d KMeans.Cluster(...) got %v after permutation, want %v", i, classes2.Index, test.partitions)
		}
	}
}

