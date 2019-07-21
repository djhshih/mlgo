package cluster

import (
	"sort"
)

type KMedians struct {
	KMeans
}

func NewKMedians(X Matrix, metric MetricOp) *KMedians {
	return &KMedians{ KMeans: *NewKMeans(X, metric) }
}

// Cluster runs the k-medians algorithm once with random initialization
// Returns the classification information
// N.B. Must explicitly override KMeans.Cluster s.t. KMedians.maximization is called
// instead of KMeans.maximization.
func (c *KMedians) Cluster(k int) (classes *Classes) {

	if c.X == nil || k >= c.Len() {
		return
	}

	c.K = k
	c.initialize()
	i := 0
	for !c.expectation() && (c.MaxIter == 0 || i < c.MaxIter) {
		c.maximization()
		i++
	}
	if i == 0 {
		// convergence is achieved right after initialization...
		// run maximization at least once to calculate cost
		c.maximization()
	}

	// copy classifcation information
	classes = &Classes{
		make([]int, c.Len()), k, c.Cost}
	copy(classes.Index, c.Clusters)

	return
}

// Override KMeans.maximization
// Calculate the median instead of mean;
// total absolute deviation instead of total sum of squares
func (c *KMedians) maximization() {
	// move cluster centroid_ii to the median
	move := func(ii int, chCost chan float64) {
		center := c.Centers[ii]
		// hold coordinate of each dimension for each member
		// members is a dimension by member matrix
		members := make(Matrix, len(center))
		// initialize to hold the maximum possible of elements
		for j, _ := range center {
			members[j] = make(Vector, len(c.Index))
		}

		// gather all member data points
		n := 0
		memberIdx := make([]int, len(c.Clusters))
		for i, class := range c.Clusters {
			if class == ii {
				for j, _ := range center {
					members[j][n] = c.X[ c.Index[i] ][j]
				}
				memberIdx[n] = c.Index[i]
				n++
			}
		}
		memberIdx = memberIdx[:n]

		// compute center as median of each data dimension
		for j, _ := range center {
			// find median
			center[j] = median(members[j][:n])
		}

		// compute cost
		cost := 0.0
		for _, i := range memberIdx {
			cost += c.Metric(center, c.X[i])
		}

		c.Errors[ii] = cost
		chCost <- cost;
	}

	// process cluster center concurrently
	ch := make(chan float64)
	for ii, _ := range c.Centers {
		go move(ii, ch)
	}

	// collect results
	J := 0.0
	for ii := 0; ii < len(c.Centers); ii++ {
		J += <-ch;
	}
	c.Cost = J / float64( c.Len() )
}

// find median
// side-effect: x becomes sorted
func median(x Vector) (med float64) {
	sort.Float64s(x)
	n := len(x)

	// calculate median
	if n % 2 == 0 {
		i := n/2
		med = (x[i] + x[i-1]) / 2
	} else {
		med = x[n/2]
	}

	/*
	// calculate total absolute deviation
	for _, z := range x {
		tad += math.Fabs( z - med )
	}
	*/

	return
}

