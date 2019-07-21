package cluster

// UnionFind implementing quick-union and find with path compression.
// NB  New sets are created for each union
// TODO Keep track of cluster sizes and use union-by-rank?

type UnionFind struct {
	// parent index array
	Parent []int
	// index for new set
	newIndex int
}

func NewUnionFind(n int) *UnionFind {
	x := &UnionFind{
		// allocate space for n singletons, and n-1 new sets
		Parent: make([]int, 2*n - 1),
		// index for new set starts after all singletons
		newIndex: n,
	}

	// Initialize index array s.t. each element is in its own set
	// i.e. each element is its own root
	for i := range x.Parent {
		x.Parent[i] = i
	}

	return x
}

// Union creates a new set join i and j together
// Does not use union-by-rank
// Same memory requirement as union-by-rank, since memory is used to expand
// the Parent array instead of maintaining a set size array
func (x *UnionFind) Union(i, j int) {
	x.Parent[x.Find(i)] = x.newIndex
	x.Parent[x.Find(j)] = x.newIndex
	x.newIndex++
}

// Find finds the index of the disjoint set in which i belongs
func (x *UnionFind) Find(i int) (r int) {
	parent := x.Parent

	// find the root by traversing up tree
	for r = i; parent[r] != r; r = parent[r] {}
	// r now holds the index of the set (i.e. root)
	// second traversal using i to perform pass compression
	for ; parent[i] != r; {
		// set all elements along path to point directly to root
		i, parent[i] = parent[i], r
	}

	return
}

// Same returns whether elements i and j belong to the same disjoint set.
func (x *UnionFind) Same(i, j int) bool {
	return x.Find(i) == x.Find(j)
}

