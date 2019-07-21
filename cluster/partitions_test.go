package cluster

import "testing"

var partitionsEqualTests = []struct {
	p, q Partitions
	equal bool
}{
	{
		Partitions{},
		Partitions{},
		true,
	},
	{
		Partitions{},
		Partitions{0, 0},
		false,
	},
	{
		Partitions{0, 0, 0},
		Partitions{0, 0},
		false,
	},
	{
		Partitions{0, 0, 0, 1, 1, 1},
		Partitions{0, 0, 0, -1, -1, -1},
		true,
	},
	{
		Partitions{0, 0, 1, 1, 1, 1},
		Partitions{0, 0, 0, 1, 1, 1},
		false,
	},
	{
		Partitions{1, 1, 1, 0, 0, 0},
		Partitions{0, 0, 0, 1, 1, 1},
		true,
	},
	{
		Partitions{1, 1, 1, 1, 1, 1},
		Partitions{0, 0, 0, 1, 1, 1},
		false,
	},
	{
		Partitions{1, 0, 2, 2, 0, 0},
		Partitions{3, 1, 4, 4, 1, 1},
		true,
	},
	{
		Partitions{3, 9, 2, 3, 1, 1},
		Partitions{0, 2, 4, 0, 1, 1},
		true,
	},
}

func TestPartitionsEqual(t *testing.T) {
	for i, test := range partitionsEqualTests {
		if test.p.Equal(test.q) != test.equal {
			t.Errorf("#%d %v.Equal(%v) got %v, want %v", i, test.p, test.q, test.p.Equal(test.q), test.equal)
		}
	}
}

