package cluster

import (
	"code.google.com/p/mlgo"
	"testing"
)

var silhouetteTests = []struct {
	x Matrix
	metric MetricOp
	classes *Classes
	silhouettes Vector
}{
	{
		Matrix{
			{101, 102, 103}, {102, 103, 104}, {103, 104, 105},
			{111, 112, 113}, {112, 113, 114}, {113, 114, 115},
			{ 21,  22,  23}, { 22,  23,  24}, { 23,  24,  25},
			{ 29,  30,  31}, { 32,  33,  34}, { 33,  34,  35},
		},
		Manhattan,
		&Classes{ Index:Partitions{0, 0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 3}, K:4 },
		Vector{
			0.8636364, 0.9000000, 0.8333333, 
			0.8333333, 0.9000000, 0.8636364,
			0.8548387, 0.8928571, 0.8200000,
			0.5000000, 0.8000000, 0.7727273,
		},
	},
}

func TestSilhouettes(t *testing.T) {
	for i, test := range silhouetteTests {
		d := NewDistances(test.x, test.metric)
		sil := Silhouettes( Segregations(d, test.classes), test.classes )
		if !mlgo.Vector(test.silhouettes).Equal(mlgo.Vector(sil)) {
			t.Errorf("#%d Silhouettes(Segregations(...), ...) got %v, want %v", i, sil, test.silhouettes)
		}
	}
}

var shadowTests = []struct {
	x, centers Matrix
	metric MetricOp
	classes *Classes
	shadows Vector
}{
	{
		Matrix{
			{101, 102, 103}, {102, 103, 104}, {103, 104, 105},
			{111, 112, 113}, {112, 113, 114}, {113, 114, 115},
			{ 21,  22,  23}, { 22,  23,  24}, { 23,  24,  25},
			{ 29,  30,  31}, { 32,  33,  34}, { 33,  34,  35},
		},
		Matrix{
			{102, 103, 104},
			{112, 113, 114},
			{ 22,  23,  24},
			{ 32,  33,  34},
		},
		Manhattan,
		&Classes{ Index: Partitions{0, 0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 3}, K: 4 },
		Vector{
			0.9090909, 1.0000000, 0.8888888,
			0.8888888, 1.0000000, 0.9090909,
			0.9090909, 1.0000000, 0.8888888,
			0.5714286, 1.0000000, 0.9090909,
		},
	},
}

func TestShadows(t *testing.T) {
	for i, test := range shadowTests {
		S := SegregationsFromCenters(test.x, test.centers, test.metric)
		shadows := Silhouettes(S, test.classes)
		if !mlgo.Vector(test.shadows).Equal(mlgo.Vector(shadows)) {
			t.Errorf("#%d Silhouettes(Separations(...), ...) got %v, want %v", i, shadows, test.shadows)
		}
	}
}

var segregateTests = []struct {
	x Matrix
	metric MetricOp
	k int
}{
	{
		Matrix{
			{1, 1}, {2, 2}, {3, 3},
			{11, 11}, {12, 12}, {13, 13},
			{21, 21}, {22, 22}, {23, 23},
		},
		Manhattan,
		3,
	},
}

func TestSegregate(t *testing.T) {
	const K = 9
	for i, test := range segregateTests {
		c := NewKMeans(test.x, test.metric)
		split := SegregateByMeanSil(c, K)
		if split.K != test.k {
			t.Errorf("#%d SegregateByMeanSil(*KMeans, %d) got %d, want %d", i, K, split.K, test.k)
			t.Errorf("Output: %v", split)
		}
	}
}

var splitTests = []struct {
	x Matrix
	metric MetricOp
	k int
	cost float64
}{
	/*
	{
		Matrix{
			{1, 1}, {2, 2}, {2, 2}, {1, 1},
			{51, 51}, {53, 53}, {51, 51}, {53, 53},
			{91, 91}, {91, 91}, {90, 90}, {90, 90},
		},
		Manhattan,
		3,
	},
	*/
	{
		Matrix{
			{1, 1}, {4, 4}, {5, 5}, {2, 2},
			{53, 53}, {57, 57}, {54, 54}, {56, 56},
			{91, 91}, {92, 92}, {94, 94}, {95, 95},
		},
		Manhattan,
		3,
		0.21904761904761905,
	},
}

func TestSplit(t *testing.T) {
	const K, L = 9, 9
	for i, test := range splitTests {
		//c := NewKMeans(test.x, test.metric)
		c := NewKMedoids(test.x, test.metric, nil)
		split := SplitByMeanSplitSil(c, K, L)
		if split.K != test.k {
			t.Errorf("#%d SplitByMeanSplitSil(*KMeans, %d, %d) got %d, want %d", i, K, L, split.K, test.k)
			t.Errorf("Output: %v", split)
			t.Errorf("Classes: %v", split.Cl)
		}
	}
}

