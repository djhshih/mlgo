package cluster

type pair struct {
	key float64
	value int
}

type pairs []pair

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Less(i, j int) bool {
	return p[i].key < p[j].key
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

