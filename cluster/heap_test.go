package cluster

import (
	"testing"
	"sort"
)

type KeyValues struct {
	r []KeyValue
}

func (x KeyValues) Len() int {
	return len(x.r)
}

func (x KeyValues) Less(i, j int) bool {
	return x.r[i].Key < x.r[j].Key
}

func (x KeyValues) Swap(i, j int) {
	x.r[i], x.r[j] = x.r[j], x.r[i]
}

func (x *KeyValues) Copy(y KeyValues) {
	x.r = make([]KeyValue, len(y.r))
	copy(x.r, y.r)
}

func (x KeyValues) Min() (min KeyValue) {
	if len(x.r) > 0 {
		min = x.r[0]
		for i := 1; i < len(x.r); i++ {
			if x.r[i].Key < min.Key {
				min = x.r[i]
			}
		}
	}
	return
}

func TestHeap(t *testing.T) {
	x := KeyValues{
		[]KeyValue {
			KeyValue{4, 1},
			KeyValue{2, 2},
			KeyValue{5, 3},
			KeyValue{8, 4},
			KeyValue{9, 5},
		},
	}
	y := KeyValues{}
	y.Copy(x)

	// create heap by using array
	h := Heap{y.r}
	h.Init()

	// create heap by sequential pushes
	g := Heap{}
	for _, a := range x.r {
		g.Push(a)
	}

	// sort original array
	sort.Sort(x)

	if s := h.Search( x.Min().Value ); s != 0 {
		t.Errorf("Element with min key found at position %d in heap, expected", s, 0)
	}

	for i := 0; i < x.Len(); i++ {
		a, b, c := x.r[i].Value, h.Pop(), g.Pop()
		if a != b {
			t.Errorf("Element with the %d-th min key in heapified array = %d, expected %d", i, b, a)
		}
		if a != c {
			t.Errorf("Element with the %d-th min key in incrementally built heap = %d, expected %d", i, c, a)
		}
	}

	y.Copy(x)
	h = Heap{y.r}
	h.Init()

	// replace min value with large value
	a := KeyValue{10, 10}
	h.Update(0, a)

	if d := h.Pop(); d != x.r[1].Value {
		t.Errorf("Value of element with min key in updated heap = %d, expected %d", d, x.r[1].Value)
	}

	// replace the last leaf with the min value
	a = KeyValue{1, 11}
	h.Update(h.Len()-1, a)

	if d := h.Pop(); d != a.Value {
		t.Errorf("Value of element with min key in updated (2) heap = %d, expected %d", d, a.Value)
	}
}

