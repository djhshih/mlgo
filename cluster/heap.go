package cluster

// min binary heap
// complete binary tree
// partial order: every node a stores a value that is less than
// or equal to that of its children
type Heap struct {
	elements []KeyValue
}

type KeyValue struct {
	Key float64
	Value int
}

func (h *Heap) Init() {
	// heapify existing array
	// assumes data is already stored in array
	// worse case complexity is O(n)
	n := len(h.elements)
	for i := n/2 - 1; i >= 0; i-- {
		siftdown(h.elements, i)
	}
}

func (h *Heap) Push(x KeyValue) {
	n := len(h.elements)
	// first place the value at the end of the heap, then sift up
	h.elements = append(h.elements, x)
	// sift appended value to correct position
	siftup(h.elements, n)
}

// Pop removes and returns the minimum element
// complexity is O(log(n))
func (h *Heap) Pop() (y int) {
	n := len(h.elements)
	if n == 0 { return -1 }
	n--
	// swap min with the last value
	h.elements[0], h.elements[n] = h.elements[n], h.elements[0]

	// get the value to remove (min)
	y = h.elements[n].Value
	// re-slice to remove the last element
	h.elements = h.elements[:n]

	// sift down new root
	if n != 0 { siftdown(h.elements, 0) }

	return
}

// Remove removes the element at position i
func (h *Heap) Remove(i int) (y int) {
	n := len(h.elements)
	if i < 0 || i >= n {
		// invalid position
		return -1
	}
	n--
	// swap with last value
	h.elements[i], h.elements[n] = h.elements[n],h.elements[i]

	// pop value
	y = h.elements[n].Value
	h.elements = h.elements[:n]
	
	if n != 0 {
		// sift up if small key
		siftup(h.elements, i)
		// sift down if large key
		siftdown(h.elements, i)
	}

	return
}


// Update updates the value at the i-th position
func (h *Heap) Update(i int, x KeyValue) {
	n := len(h.elements)
	if i < 0 || i >= n {
		// invalid position
		return
	}
	
	oldKey := h.elements[i].Key
	h.elements[i] = x

	if x.Key < oldKey {
		// value became smaller: siftup
		siftup(h.elements, i)
	} else if x.Key > oldKey {
		// value became bigger: siftdown
		siftdown(h.elements, i)
	}
	// no action if key did not change
}

func (h *Heap) Len() int {
	return len(h.elements)
}

func (h *Heap) Search(value int) (i int) {
	i = -1
	for j, a := range h.elements {
		if a.Value == value {
			i = j
		}
	}
	return
}

func siftup(elements []KeyValue, i int) {
	// sift up until x's parent <= x
	// i, j are the indices of x and x's parent
	for ; i != 0; {
		// parent
		j := (i-1)/2
		if elements[j].Key <= elements[i].Key {
			break
		}
		// sift up
		elements[i], elements[j] = elements[j], elements[i]
		i = j
	}
}

func siftdown(elements []KeyValue, i int) {
	// n is the new size of the elements array
	// elements array has not been re-resized yet
	n := len(elements)

	// loop until i is a leaf
	for ; !(i >= n/2 && i < n); {
		j := 2*i + 1     // left child
		right := j + 1   // right child: 2*i + 2
		if right < n && elements[right].Key < elements[j].Key {
			// set j to lesser child
			j = right
		}
		if elements[i].Key <= elements[j].Key {
			// i is less than both its children: heap order satisified
			break
		}
		// sift down
		elements[i], elements[j] = elements[j], elements[i]
		i = j
	}
}

