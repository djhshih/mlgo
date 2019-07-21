package cluster

type Partitions []int

func (p Partitions) Len() int {
	return len(p)
}

func (p Partitions) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p Partitions) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Reassign reassigns partition labels so that partitions 
// are assigned indices (1-index) in the order of appearance.
func (p Partitions) Reassign() {
	// labels will store the new labels
	// labels[i] will store the new index for index i in p
	// by default, labels[i] == 0
	labels := make(map[int] int, len(p))

	next := 1
	for i := range p {
		if labels[ p[i] ] == 0 {
			// set new label to next available index
			labels[ p[i] ] = next
			next++
		}
		// assign new label
		p[i] = labels[ p[i] ]
	}
}

// Equal returns whether partitions p and q are equal.
func (p Partitions) Equal(q Partitions) bool {
	if len(p) != len(q) {
		return false
	}
	if len(p) == 0 {
		// both p and q are empty
		return true
	}

	// copy partitions
	a, b := make(Partitions, len(p)), make(Partitions, len(q))
	copy(a, p)
	copy(b, q)

	// re-assign partition labels to standardize
	a.Reassign()
	b.Reassign()

	// element-by-element comparison
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

