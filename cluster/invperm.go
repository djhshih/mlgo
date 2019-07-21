package cluster

// InvPerm inverses the permutation x
func InvPerm(x []int) {
	n := len(x)
	for i := 0; i < n; i++ {
		x[i] = -x[i] - 1
	}
	for m := n - 1; m >= 0; m-- {
		i, j := m, x[m]
		for ; j >= 0; j = x[j] {
			i = j
		}
		x[i] = x[-j-1]
		x[-j-1] = m
	}
}

func PermEqual(x, y[]int) bool {
	n := len(x)
	if n != len(y) {
		return false
	}
	if n == 0 {
		return true
	}
	for i := 0; i < n; i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

// Permute returns the permutation of int slice x using index p.
func Permute(x, p []int) (y []int) {
	y = make([]int, len(p))
	for i, v := range p {
		y[i] = x[v]
	}
	return
}

