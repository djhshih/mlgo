package cluster

import (
	"code.google.com/p/mlgo"
	"sort"
)

// CoordinatesSetEqual returns whether the a and b contain the same set of coordinates.
// Each row of a and b is a tuple of coordinates.
func CoordinatesSetEqual (X, Y Matrix) bool {

	// declare mlgo.Matrix, which has all the associated methods
	var A, B mlgo.Matrix
	// copy matrices, since they will be sorted
	A = mlgo.CopyMatrix(X)
	B = mlgo.CopyMatrix(Y)

	// sort the rows to standardize the order
	// i.e. eliminate order difference, since order does not matter for sets
	sort.Sort(A)
	sort.Sort(B)

	// check if sorted matrices are equal
	return A.Equal(B)
}

