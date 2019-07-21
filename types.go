package mlgo

import (
	"math"
)

type Vector []float64

// this definition creates problem in client program defining Matrix and Vector
// i.e. mlgo.Vector != Vector
//type Matrix []Vector
type Matrix [][]float64

const (
	MaxValue = math.MaxFloat64
)

func Range(start, end int) (x []int) {
	size := end - start
	x = make([]int, size)
	v := start
	for i := 0; i < size; i++ {
		x[i] = v
		v++
	}
	return
}

func NewVector(m int, v float64) (x Vector) {
	x = make(Vector, m)
	for i := 0; i < m; i++ {
		x[i] = v
	}
	return
}

func (x Vector) Reordered(index []int) (y Vector) {
	y = make(Vector, len(index))
	for i, v := range index {
		y[i] = x[v]
	}
	return
}

func (x Vector) Mean() (m float64) {
	for i := 0; i < len(x); i++ {
		m += x[i]
	}
	m /= float64(len(x))
	return
}

func (x Vector) Summarize() (mean, variance float64) {
	var stats Summary
	for _, v := range x {
		// accumulate statistics
		stats.Add(v)
	}

	mean, variance = stats.Mean, stats.VarP()
	return
}

func (x Vector) Equal(y Vector) bool {
	const epsilon = 1e-6
	for i := range x {
		if !EssentiallyEqual(x[i], y[i], epsilon)  {
			return false
		}
	}
	return true
}

func (X Matrix) Summarize() (means, variances Vector) {
	m := len(X)
	if m < 2 { return }

	n := len(X[0])
	stats := make([]Summary, n)

	means, variances = make(Vector, n), make(Vector, n)

	for i := 0; i < m; i++ {
		// accumulate statistics for each feature
		for j, x := range X[i] {
			stats[j].Add(x)
		}
	}

	for j, _ := range stats {
		means[j] = stats[j].Mean
		variances[j] = stats[j].VarP()
	}

	return
}

func (X Matrix) Len() int {
	return len(X)
}

// Less returns whether row i is lexicographically less than row j
func (X Matrix) Less(i, j int) bool {
	for k := range X[i] {
		if X[i][k] < X[j][k] {
			return true
		} else if X[i][k] > X[j][k] {
			return false
		}
		// otherwise, X[i][k] == X[j][k] and continue to next position
	}
	// all positions are equal
	return false
}

func (X Matrix) Swap(i, j int) {
	X[i], X[j] = X[j], X[i]
}


// CopyMatrix returns a Matrix filled with content from Y
func CopyMatrix(Y [][]float64) (X Matrix) {
	X = make(Matrix, len(Y))
	for i := range Y {
		X[i] = make(Vector, len(Y[i]))
		copy(X[i], Y[i])
	}
	return
}

// Copied returns a copy of the Matrix.
func (X Matrix) Copied() (Y Matrix) {
	Y = make(Matrix, len(X))
	for i := range X {
		Y[i] = make(Vector, len(X[i]))
		copy(Y[i], X[i])
	}
	return
}

func (X Matrix) Slice(idx []int) (Y Matrix) {
	Y = make(Matrix, len(idx))
	for i, ii := range idx  {
		Y[i] = make(Vector, len(X[ii]))
		copy(Y[i], X[ii])
	}
	return
}

func (X Matrix) Equal(Y Matrix) bool {
	const epsilon = 1e-6
	for i := range X {
		for j := range X[i] {
			if !EssentiallyEqual(X[i][j], Y[i][j], epsilon) {
				return false
			}
		}
	}
	return true
}

