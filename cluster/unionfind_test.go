package cluster

import (
	"testing"
)

var unionFindTests = []struct {
	n int
	unions [][]int
	i, j int
	same bool
}{
	{
		4,
		[][]int{ {0, 1}, {1, 2} },
		0, 2,
		true,
	},
	{
		4,
		[][]int{ {0, 1}, {1, 2} },
		1, 3,
		false,
	},
	{
		6,
		[][]int{ {0, 5}, {1, 6}, {2, 7}, {4, 8} },
		5, 4,
		true,
	},
	{
		6,
		[][]int{ {0, 5}, {1, 6}, {2, 7}, {4, 8} },
		3, 1,
		false,
	},
}

func TestUnionFind(t *testing.T) {
	for i, test := range unionFindTests {
		uf := NewUnionFind(test.n)
		for _, union := range test.unions {
			uf.Union(union[0], union[1])
		}
		if uf.Same(test.i, test.j) != test.same {
			t.Errorf("#%d UnionFind.Same(%d, %d) got %v, want %v", i, test.i, test.j, uf.Same(test.i, test.j), test.same)
			t.Errorf("unions [%d]: %v", test.n, test.unions)
			t.Errorf("UnionFind.Parent: %v", uf.Parent)
		}
	}
}

