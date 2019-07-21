package mlgo

import (
	"math"
)

type Summary struct {
	Mean, N, devsq, Min, Max float64
}

// Add accumulates running statistics for calculating variance and
// standard deviation using the Welford method (1962)
func (s *Summary) Add(x float64) {
	if s.N > 0 {
		if x < s.Min { s.Min = x }
		if x > s.Max { s.Max = x }
	} else {
		s.Min, s.Max = x, x
	}

	s.N++
	t := x - s.Mean
	s.Mean += t / s.N
	s.devsq += t * (x - s.Mean)

}

func (s *Summary) AddValues(x []float64) {
	for _, z := range x {
		s.Add(z)
	}
}

// Var returns the sample variance
func (s *Summary) Var() (v float64) {
	if s.N > 2 {
		v = s.devsq / (s.N - 1)
	}
	return
}

// Sd returns the sample standard deviation
func (s *Summary) Sd() (v float64) {
	if s.N > 2 {
		v = math.Sqrt( s.devsq / (s.N-1) )
	}
	return
}

// VarP returns the population variance
func (s *Summary) VarP() (v float64) {
	if s.N > 1 {
		v = s.devsq / s.N
	}
	return
}

// SdP returns the population standard deviation
func (s *Summary) SdP() (v float64) {
	if s.N > 1 {
		v = math.Sqrt( s.devsq / s.N )
	}
	return
}

// Range returns the range of the data
func (s *Summary) Range() (r float64) {
	if s.N > 1 {
		r = s.Max - s.Min
	}
	return
}

