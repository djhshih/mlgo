package cluster

import (
	"code.google.com/p/mlgo"
	"math"
)

// Validation measures

// Silhouette considerations
// use different linkage methods to pre-calculate distance between an element and a cluster
// pre-calculate these linkage distances for efficiency
// use update formula?

// Segregations return a matrix of distances between data points and clusters
func Segregations(distances *Distances, classes *Classes) (S Matrix) {
	// each row of x is considered one data point
	m := distances.Len()

	// allocate space
	S = make(Matrix, m)
	for i := 0; i < m; i++ {
		S[i] = make(Vector, classes.K)
	}

	index := classes.Index

	// determine cluster sizes
	sizes := classes.Sizes()

	// calculate the average distances from data point i to data points, for each cluster
	// TODO option to aggregate by median/max/min instead of mean?
	for i := 0; i < m; i++ {
		// accumulate sum
		for j := 0; j < m; j++ {
			S[i][ index[j] ] += distances.Get(i, j)
		}
		// derive mean via division by cluster sizes
		for jj := 0; jj < classes.K; jj++ {
			if sizes[jj] > 1 {
				S[i][jj] /= float64(sizes[jj])
			} else if sizes[jj] == 0 {
				S[i][jj] = math.Inf(1)
			}
		}
		// correct mean for own cluster (divide by size-1 instead of size)
		// element in singleton cluster has 0 distance to itself
		c := index[i]
		size := float64(sizes[c])
		if size > 1 {
			S[i][c] *= size / (size - 1)
		}
	}
	return
}

// SegregationsFromCenters return a matrix of distances between data points and cluster centers
func SegregationsFromCenters(X, centers Matrix, metric MetricOp) (S Matrix) {
	// each row of x is considered one data point
	m := len(X)
	k := len(centers)

	// allocate space
	S = make(Matrix, m)
	for i := 0; i < m; i++ {
		S[i] = make(Vector, k)
	}

	// calculate distance from data point i to center of cluster j
	for i := 0; i < m; i++ {
		for j := 0; j < k; j++ {
			S[i][j] = metric(X[i], centers[j])
		}
	}
	return
}

// Silhouettes returns a vector of silhouettes for data points.
// If S is a matrix of average distances from each elements to other elements in each cluster, then the returned values are conventionally considered as silhouettes.
// If S is a matrix of distances from each element to each cluster center, then the returned values are can be considered as shadows.
// TODO special case: silhouette is not defined for two singleton clusters
// TODO faithful calculation of "shadow" as defined by Friedrich Leisch (average two nearest centroids for 'b')
func Silhouettes(S Matrix, classes *Classes) (s Vector)  {
	m := len(S)
	k := len(S[0])

	s = make(Vector, m)

	index := classes.Index
	sizes := classes.Sizes()

	// calculate silouettes
	for i := 0; i < m; i++ {
		c := index[i]
		if sizes[c] == 1 {
			// element is in a singleton class: silhouette is 0 by definition
			//s[i] = 0
			continue
		}
		// distance to own cluster
		a := S[i][c]
		// distance to nearest cluster
		b := math.Inf(1)
		for j := 0; j < k; j++ {
			if j != c && S[i][j] < b {
				b = S[i][j]
			}
		}
		if b < math.Inf(1) {
			max := a
			if a < b {
				max = b
			}
			s[i] = (b - a) / max
		} else {
			// no other cluster is available: set silhouette to 0
			//s[i] = 0
		}
	}
	return
}


type Split struct {
	K int
	Cl *Classes
	Cost float64
}

type Segregator interface {
	Clusterer
	Segregations(classes *Classes) Matrix
}

// TODO Do not count the silhouette of singleton clusters in the average?
func SegregateByMeanSil(seg Segregator, K int) (s Split) {
	m := seg.Len()

	// silhouette can only be calculated for 2 <= k <= m - 1

	if K <= 0 || K > m - 1 {
		K = m - 1
	}

	// maximize average silhouette
	avgSil := -1.0
	optK := 0
	var optClasses *Classes
	for k := 2; k <= K; k++ {
		classes := seg.Cluster(k)
		sil := Silhouettes(seg.Segregations(classes), classes)
		t := mlgo.Vector(sil).Mean()
		if t > avgSil {
			avgSil = t
			optK = k
			optClasses = classes
		}
	}

	s.K = optK
	s.Cost = 1 - avgSil
	s.Cl = optClasses
	return
}

type Splitter interface {
	Segregator
	Subset(index []int) Splitter
}

//TODO median split silhouette

// K is the maximum number of clusters.
// L is the maximum number of children clusters for any cluster.
func SplitByMeanSplitSil(splitter Splitter, K, L int) (s Split) {
	m := splitter.Len()

	// average split silhouette can be only be calculated for 1 <= k <= m/3
	// if k > m/3, at least one cluster would have < 3 elements
	// each cluster needs >= 3 elements to be further split into at least 2 clusters 
	//  for silhouette calculation

	if K <= 0 || K > m / 3 {
		K = m / 3
	}

	// minimize the mean split silhouette
	avgSplitSil := math.Inf(1)
	optK := 0
	var optClasses *Classes
	for k := 1; k <= K; k++ {
		splitSil := make(Vector, k)
		classes := splitter.Cluster(k)
		partitions := classes.Partitions()
		n := 0
		for kk := 0; kk < classes.K; kk++ {
			clustSplit := SegregateByMeanSil(splitter.Subset(partitions[kk]), L)
			if clustSplit.K > 0 {
				// cluster could be split further into children clusters
				splitSil[n] = 1 - clustSplit.Cost
				n++
			}
		}
		// remove empty elements at end to account for clusters that could be not split further
		splitSil = splitSil[:n]
		t := mlgo.Vector(splitSil).Mean()
		//fmt.Println(k, t, splitSil, classes)
		if t < avgSplitSil {
			avgSplitSil = t
			optK = k
			optClasses = classes
		}
	}

	s.K = optK
	s.Cost = avgSplitSil
	s.Cl = optClasses
	return
}

