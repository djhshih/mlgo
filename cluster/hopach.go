package cluster

type Hopacher interface {
	Clusterer
	Subset(index []int) Hopacher
	// Heterogeneity of a given partitioning
	Heterogeneity(classes *Classes) float64
	// Sort elements in some order
	Sort()
}

type Hopach struct {
	Base Hopacher

	maxLevel, maxK, maxL int
	// implemented parameters
	// clusters = best
	// coll = seq
	// newmed = nn
	// mss = med
	// initord = co
	// ord = neighbour
}

func NewHopach(base Hopacher) *Hopach {
	return &Hopach{
		Base: base,
		maxLevel: 15,
		maxK: 9,
		maxL: 9,
	}
}

func (h *Hopach) Hierarchize() Linkages {

	return nil
}

