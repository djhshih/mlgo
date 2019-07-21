package cluster

import (
	"testing"
)

var invPermTests = []struct {
	in, ans []int
}{
	{ []int{5, 4, 3, 2, 1, 0}, []int{5, 4, 3, 2, 1, 0} },
	{ []int{0, 1, 2, 3, 4, 5}, []int{0, 1, 2, 3, 4, 5} },
	{ []int{0, 2, 4, 1, 3, 5}, []int{0, 3, 1, 4, 2, 5} },
	{ []int{3, 1, 4, 0, 5, 2}, []int{3, 1, 5, 0, 2, 4} },
}

func TestInvPerm(t *testing.T) {
	for i, test := range invPermTests {
		out := make([]int, len(test.in))
		copy(out, test.in)
		InvPerm(out)
		if !PermEqual(out, test.ans) {
			t.Errorf("#%d after InvPerm(%v), got %v, want %v", i, test.in, out, test.ans)
		}
	}
}

