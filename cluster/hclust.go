package cluster

const (
	single_linkage = iota
	complete_linkage
	average_linkage
	mcquitty_linkage
	median_linkage
	centroid_linkage
	ward_linkage
)

type HClusters struct {
	// Data points [m x n]
	X Matrix
	// Distance metric
	Metric MetricOp
	// number of clusters
	K int
	// linkage method
	Method int
	// Distances between data points [m x m]
	D *Distances
	// Step-wise dendrogram
	Dendrogram Linkages
	// cluster center assignment index
	Index []int
	// cost
	Cost float64
	// indices of active elements
	actives ActiveSet
}

// CutTree cuts the hierarchical cluster tree to generate K clusters.
func (c *HClusters) CutTree(K int) {
	if c.Dendrogram == nil { return }

	if K == 0 {
		// by default, leave each element in its own cluster
		K = len(c.X)
	}
	c.K = K

	m := len(c.X)
	uf := NewUnionFind(m)

	// Starting with each element in its own cluster
	// after each merge in the stepwise dendrogram, one less cluster remains
	// therefore, (m - K) merges will occur
	for i := 0; i < m-K; i++ {
		linkage := c.Dendrogram[i]
		uf.Union(linkage.First, linkage.Second)
	}
	
	// Now that all the merges have been done, determine the cluster index
	c.Index = make([]int, m)
	for i := 0; i < m; i++ {
		c.Index[i] = uf.Find(i)
	}
}

// CutTreeHeight cuts the hierarchical cluster tree to specified height.
func (c *HClusters) CutTreeHeight(height float64) {
	if c.Dendrogram == nil { return }

	m := len(c.X)
	uf := NewUnionFind(m)

	// Starting with each element in its own cluster
	// after each merge in the stepwise dendrogram, one less cluster remains
	// therefore, (m - k) merges will occur
	for i := 0; i < m-1; i++ {
		linkage := c.Dendrogram[i]
		if linkage.Distance > height {
			// i merges have occured: m - i clusters remain
			c.K = m - i
			break
		}
		uf.Union(linkage.First, linkage.Second)
	}
	
	// Now that all the merges have been done, determine the cluster index
	c.Index = make([]int, m)
	for i := 0; i < m; i++ {
		c.Index[i] = uf.Find(i)
	}
}

// linkage
type Linkage struct {
	First, Second int
	Distance float64
}

type Linkages []Linkage

func (x Linkages) Len() int {
	return len(x)
}

func (x Linkages) Less(i, j int) bool {
	return x[i].Distance < x[j].Distance
}

func (x Linkages) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

