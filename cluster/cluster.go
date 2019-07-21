package cluster

import (
	"code.google.com/p/mlgo"
)

type Vector mlgo.Vector
type Matrix mlgo.Matrix

const maxValue = mlgo.MaxValue

type Classes struct {
	// classification index
	Index Partitions
	K int
	Cost float64
}

func (c *Classes) Sizes() []int {
	sizes := make([]int, c.K)
	index := c.Index
	m := len(index)
	for i := 0; i < m; i++ {
		sizes[index[i]]++
	}
	return sizes
}

// Partitions return an array of partition element arrays
func (c *Classes) Partitions() [][]int {
	// allocate space
	sizes := c.Sizes()
	p := make([][]int, c.K)
	for i := 0; i < c.K; i++ {
		p[i] = make([]int, sizes[i])
	}
	counters := make([]int, c.K)
	// determine partitions
	m := len(c.Index)
	for i := 0; i < m; i++ {
		cl := c.Index[i]
		p[cl][counters[cl]] = i
		counters[cl]++
	}
	return p
}

type Clusterer interface {
	// Cluster clusters data points into k clusters.
	Cluster(k int) *Classes
	Len() int
}

type Subclusterer interface {
	Clusterer
	// Subcluster clusters a subset of data points specified by idx into k clusters.
	Subcluster(k int, idx []int) *Classes
}

type Hierarchizer interface {
	// Hierarchize organizes data clusters in a dendrogram
	Hierarchize() Linkages
}

// FIXME There should be multiple instances of Clusterer
// FindClusters runs the clustering algorithm for the specified number of repeats.
func FindClusters(c Clusterer, k int, repeats int) (classes *Classes) {
	// repeat clustering concurrently
	ch := make(chan *Classes)
	for i := 0; i < repeats; i++ {
		go func() {
			ch <- c.Cluster(k)
		}()
	}

	// determine best clustering
	minCost := maxValue
	for i := 0; i < repeats; i++ {
		if cl := <-ch; cl.Cost < minCost {
			classes = cl
			minCost = cl.Cost
		}
	}
	return
}

