package cluster

import (
	"sort"
)

// TODO make KMedoids use Distances class
// FIXME KMedoids may initialize to 'duplicate' data points when there are ties,
//       which will result in fewer than k clusters

type KMedoids struct {
	KMeans
}

func NewKMedoids(X Matrix, metric MetricOp, distances *Distances) *KMedoids {
	if distances == nil {
		distances = NewDistances(X, metric)
	}
	c := &KMedoids{
		KMeans: *NewKMeans(X, metric),
	}
	c.D = distances
	return c
}

// Cluster runs the k-medoids algorithm.
// Returns the classification information.
func (c *KMedoids) Cluster(k int) (classes *Classes) {
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
		make([]int, c.Len()), k, c.Cost }
	copy(classes.Index, c.Clusters)

	return
}

func (c *KMedoids) Subset(index []int) Splitter {
	D := c.D.Subset(index)
	return &KMedoids{
		KMeans: KMeans{X: c.X, Metric: c.Metric, Index: index, D: D},
	}
}


// Initialize the medoids by choosing the most central k data points.
func (c *KMedoids) initialize() {
	m := c.Len()

	// calculate normalized distances
	normalized := make(Matrix, m)
	for i := 0; i < m; i++ {
		normalized[i] = make(Vector, m)
		sum := 0.0
		for j := 0; j < m; j++ {
			d := c.D.Get(i, j)
			normalized[i][j] = d
			sum += d
		}
		for j := 0; j < m; j++ {
			normalized[i][j] /= sum
		}
	}

	// sum the normalized distances across all rows
	p := make(pairs, m)
	for i, _ := range(normalized) {
		p[i].value = c.Index[i]
		for _, x := range normalized[i] {
			p[i].key += x
		}
	}

	// sort the summed normalized distances
	sort.Sort(p)

	c.Clusters = make([]int, c.Len())

	// initialize centers
	c.Centers, c.Errors = make(Matrix, c.K), make(Vector, c.K)
	for k, _ := range c.Centers {
		// use the first k data points sorted by summed normalized distances
		x := c.X[ p[k].value ]
		c.Centers[k] = make(Vector, len(x))
		copy(c.Centers[k], x)
	}
}

// Maximization step: Swap medoid with another data point in the cluster
// s.t. total distance to the new medoid is minimized.
func (c *KMedoids) maximization() {
	// swap medoid
	swap := func(ii int, chCost chan float64) {
		center := c.Centers[ii]

		// gather members
		n := 0
		memberIdx := make([]int, len(c.Clusters))
		for i, class := range c.Clusters {
			if class == ii {
				memberIdx[n] = i
				n++
			}
		}

		if n == 0 {
			// medoid has no member: terminate with 0 cost
			chCost <- 0
			return
		}
		memberIdx = memberIdx[:n]

		// calculate total distances for each member
		totalDistances := make(Vector, n)
		n = 0
		for _, i := range memberIdx {
			for _, j := range memberIdx {
				totalDistances[n] += c.D.Get(i, j)
			}
			n++
		}

		// find the member with the minimum total distance
		// set this as the new center
		newCenter, min := memberIdx[0], totalDistances[0]
		for i, d := range totalDistances {
			if d < min {
				newCenter, min = memberIdx[i], d
			}
		}
		copy(center, c.X[newCenter])

		// use the minimum total distance as the cost
		c.Errors[ii] = min
		chCost <- min
	}

	// process cluster center concurrently
	ch := make(chan float64)
	for ii, _ := range c.Centers {
		go swap(ii, ch)
	}

	// collect results
	J := 0.0
	for ii := 0; ii < len(c.Centers); ii++ {
		J += <-ch;
	}
	c.Cost = J / float64( len(c.X) )
}

