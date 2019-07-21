package cluster

import (
	"testing"
)

func TestActiveSet(t *testing.T) {
	n := 10
	activeSet := NewActiveSet(n)
	for i := 0; i < n; i++ {
		if !activeSet.Contains(i) {
			t.Errorf("Newly initialized activeSet is corrupt; activeSet = %v, should contain %d", activeSet, i)
		}
	}

	removeIdxSet := []int{0, 2, 6, 9}
	for _, i := range removeIdxSet {
		activeSet.Remove(i)
		if activeSet.Contains(i) {
			t.Errorf("Removed element %d still exists in activeSet", i)
		}
	}
	if activeSet.first == 0 {
		t.Errorf("activeSet.First points to a removed element")
	}

	count := 0
	for i := activeSet.Begin(); i != activeSet.End(); i = activeSet.Next(i) {
		count++
		if count > n {
			t.Errorf("activeSet has more elements than capacity (%d); possible circular reference", n)
		}
	}

	n2 := n - len(removeIdxSet)
	if count != n2 {
		t.Errorf("activeSet after removal does not contain the expected number of elements; count = %d, expected %d", count, n2)
	}
	if activeSet.Len() != n2 {
		t.Errorf("activeSet.Len() == %d, expected %d", activeSet.Len(), n2)
	}

	// At this point, the elements are: [1, 3, 4, 5, 7, 8]
	if activeSet.Get(3) != 5 {
		t.Errorf("activeSet.Get(3) == %d, expected 5", activeSet.Get(3))
	}
	if activeSet.Get(7) != 3 {
		t.Errorf("activeSet.Get(7) == %d, expected 3", activeSet.Get(7))
	}
	if activeSet.Get(-4) != 4 {
		t.Errorf("activeSet.Get(-4) == %d, expected 4", activeSet.Get(4))
	}

}

