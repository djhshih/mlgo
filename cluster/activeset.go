package cluster

// array-based integer singly linked list
// can remove nodes but cannot add nodes
type ActiveSet struct {
	first int
	values []int
	capacity, size int
}

func NewActiveSet(n int) (l ActiveSet) {
	l = ActiveSet{
		first: 0,
		values: make([]int, n),
		capacity: n,
		size: n,
	}
	l.Init()
	return
}

func (l *ActiveSet) Init() {
	for i := 0; i < l.capacity; i++ {
		// each element stores the next value
		l.values[i] = i+1
	}
}

func (l *ActiveSet) Next(curr int) int {
	return l.values[curr]
}

func (l *ActiveSet) Begin() int {
	return l.first
}

func (l *ActiveSet) End() int {
	return l.capacity
}

func (l *ActiveSet) Remove(i int) {
	if i < 0 || i >= l.capacity || l.values[i] == -1 { return }

	if i == l.first {
		// re-assign first to the next node
		l.first = l.values[i]
	} else if i > 0 {
		// find the previous element
		prev := 0
		for j := l.first; j < i; j = l.values[j] {
			prev = j
		}
		// link the previous element to the next node
		l.values[prev] = l.values[i]
	}
	// explicitly mark element as deleted
	l.values[i] = -1
	l.size--
}

func (l *ActiveSet) Contains(i int) bool {
	if i < l.capacity && l.values[i] != -1 {
		return true
	}
	return false
}

func (l *ActiveSet) Len() int {
	return l.size
}

// Get returns the jth active element in the set, with boundary wrapping.
func (l *ActiveSet) Get(j int) int {
	// wrap overflow values
	// N.B. In Go, (-a)%b == -(a%b)
	j %= l.size
	// wrap negative values
	if j < 0 {
		j += l.size
	}

	// loop to the jth active element
	i := l.Begin()
	for k := 0; k < j; k++ {
		i = l.Next(i)
	}
	return i
}

