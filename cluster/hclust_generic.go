package cluster

//TODO  Add handling of different metrics

import "fmt"

// node
type node struct {
	// index of nearest neighbour
	nearest int
	// distance to nearest neighbour
	minDistance float64
	// size of cluster headed by node
	size int
}

// Generic hierarchical clustering using Mullner's algorithm
type HClustersGeneric struct {
	HClusters
	// one node for each data point
	nodes []node
	// priority queue (key: minimum distances; value: node index)
	priority Heap
}

func NewHClustersGeneric(X Matrix, metric MetricOp, method int, d *Distances) *HClustersGeneric {
	if d == nil {
		d = NewDistances(X, metric)
	}
	return &HClustersGeneric{
		HClusters: HClusters{
			X: X,
			Metric: metric,
			Method: method,
			D: d,
		},
	}
}

func (c *HClustersGeneric) Cluster(k int) (classes *Classes) {
	if c.D == nil { return }

	c.initialize()

	c.cluster()

	c.CutTree(k)

	// copy classification information
	classes = &Classes{
		make([]int, len(c.X)), k, c.Cost }
	copy(classes.Index, c.Index)

	return
}

// assume initialization has been run
func (c *HClustersGeneric) cluster() {
	m := len(c.X)

	// NB In updating nearest neighbour, only the node with the smaller index
	//    has the correct information in a pair of nearest neighbour,

	// (m - 1) merges will occur in main loop
	for i := 0; i < m-1; i++ {

		// Choose a pair of nodes to merge

		// get next nearest pair of nearest neighbours
		a := c.priority.Pop()
		nodeA := c.nodes[a]
		b, distance := nodeA.nearest, nodeA.minDistance

		// Re-calculate nearest neighbour, if necessary
		for ; distance > c.D.rep[a][b]; {
			// find new nearest neighbour for a
			x := c.actives.Next(a)
			min, minIdx := c.D.rep[a][x], x
			for x = c.actives.Next(x); x < m; x = c.actives.Next(x) {
				if c.D.rep[a][x] < min {
					min = c.D.rep[a][x]
				}
			}
			// update priority queue and node
			c.priority.Push( KeyValue{Value:a, Key:min} )
			nodeA.nearest = minIdx

			// get next nearest pair of nearest neighbours again
			a = c.priority.Pop()
			nodeA = c.nodes[a]
			b, distance = nodeA.nearest, nodeA.minDistance
		}
		// element a with min minDistance has been popped from the priority queue

		// Merge the pair of nearest nodes
		// insert into dendrogram
		c.Dendrogram[i] = Linkage{ First:a, Second:b, Distance:distance }
		// use b as the index for the new node
		nodeB := c.nodes[b]
		nodeB.size += nodeA.size
		// mark node a as inactive
		c.actives.Remove(a)


		// FIXME: a was set a removed node, why?

		// Update the distance matrix
		fmt.Println("loop")
		for x := c.actives.Begin(); x < m; x = c.actives.Next(x) {
			if x != b {
				// TODO use different update rule for different linkage methods
				sizeA, sizeB := float64(nodeA.size), float64(nodeB.size)
				var d float64
				switch c.Method {
					case single_linkage:
						fmt.Println(a, b, x, c.actives)
						if c.D.rep[a][x] <= c.D.rep[b][x] {
							d = c.D.rep[a][x]
						} else {
							d = c.D.rep[b][x]
						}
					case complete_linkage:
						if c.D.rep[a][x] >= c.D.rep[b][x] {
							d = c.D.rep[a][x]
						} else {
							d = c.D.rep[b][x]
						}
					case average_linkage:
						d = (sizeA * c.D.rep[a][x] + sizeB * c.D.rep[b][x]) / (sizeA + sizeB)
					case mcquitty_linkage:
						d = 0
					case median_linkage:
						d = 0
					case centroid_linkage:
						d = 0
					case ward_linkage:
						d = 0
				}
				c.D.rep[b][x] = d
				c.D.rep[x][b] = d
			}
		}

		// Update candidates for nearest neighour,
		// to be corrected in the next iteration, if necessary
		for x := c.actives.Begin(); x < a; x = c.actives.Next(x) {
			if c.nodes[x].nearest == a {
				// a was the nearest neighbour for x; but a has been removed
				// set nearest neigbour to b temporarily
				// the search for the true nearest neighbour is deferred until it is needed
				// in future iteration of the main loop
				c.nodes[x].nearest = b
			}
		}
		
		// Check if other nodes now have b as the nearest node
		// Since the current nearest neighbour may be inaccurate...
		for x := c.actives.Begin(); x < b; x = c.actives.Next(x) {
			if c.D.rep[x][b] < c.nodes[x].minDistance {
				// b is now the nearest neighbour for x
				c.nodes[x].nearest = b
				// preserve a lower bound for minDistance
				d := c.D.rep[x][b]
				c.nodes[x].minDistance = d
				// update priority queue: bottle neck in worst case time complexity
				i := c.priority.Search(x)
				c.priority.Update(i, KeyValue{d, x})
			}
		}

		// Update nearest neighbour for node b
		min, minIdx := nodeB.minDistance, nodeB.nearest
		for x := c.actives.Next(b); x < m; x = c.actives.Next(x) {
			if c.D.rep[b][x] < min {
				min, minIdx = c.D.rep[b][x], x
			}
		}
		c.nodes[b].minDistance, c.nodes[b].nearest = min, minIdx
		// update priority queue
		i := c.priority.Search(b)
		c.priority.Update(i, KeyValue{min, minIdx})

	}
}

func (c *HClustersGeneric) initialize() {
	m := len(c.X)

	c.Dendrogram = make([]Linkage, m-1)

	c.nodes = make([]node, m)

	// set of indices of active nodes
	c.actives = NewActiveSet(m)

	// Generate the list of nearest neighbours
	// iterate from the first to the penultimate node
	for i := c.actives.Begin(); i < m-1; i = c.actives.Next(i) {
		// set nearest neighbour to the next node
		j := i+1
		min, minIdx := c.D.rep[i][j], j
		// check later nodes
		for j = i+2; j < m; j++ {
			if c.D.rep[i][j] < min {
				min = c.D.rep[i][j]
			}
		}
		c.nodes[i] = node{nearest:minIdx, minDistance:min, size:1}
		// not necessary to create reciprocal relationship
		//nodes[minIdx] = node{i, min, 1}
	}
	
	// Create priority queue

	// create array with len and capacity m
	minDistances := make([]KeyValue, m, m)
	for i := 0; i < m; i++ {
		minDistances[i] = KeyValue{ Key:c.nodes[i].minDistance, Value:i }
	}
	// heapify array and store heap as class member
	c.priority = Heap{ minDistances }
	c.priority.Init()

}

