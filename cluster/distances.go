package cluster

import (
	"code.google.com/p/mlgo"
)

//TODO make Distances memory efficient

type Distances struct {
	// distances between pairs of data points
	rep Matrix
	metric MetricOp
	index []int
}

func NewDistances(X Matrix, metric MetricOp) (d *Distances)  {
	// each row of X is considered one data point
	m := len(X)

	// allocate space
	D := make(Matrix, m)
	for i := 0; i < m; i++ {
		D[i] = make(Vector, m)
	}

	// calculate distances for lower and upper triangles together
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			d := metric(X[i], X[j])
			D[i][j], D[j][i] = d, d
		}
	}

	d = &Distances{ rep: D, metric: metric, index: mlgo.Range(0, m) }

	return
}

func (d *Distances) Len() int {
	return len(d.index)
}

func (d *Distances) Subset(index []int) *Distances {
	return &Distances{ rep: d.rep, metric: d.metric, index: index }
}

func (d *Distances) Get(i, j int) float64 {
	return d.rep[d.index[i]][d.index[j]]
}

