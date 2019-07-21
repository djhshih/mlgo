package mlgo

import (
	"testing"
	"fmt"
)

func TestSummary(t *testing.T) {
	s := Summary{}
	s.AddValues([]float64{1, 2, 3, 4})
	fmt.Println(s.Mean, s.N, s.Var(), s.VarP(), s.Min, s.Max, s.Range())
}

