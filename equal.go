package mlgo

import "math"

func ApproximatelyEqual(a, b, epsilon float64) bool {
	var t float64
	if math.Abs(a) < math.Abs(b) {
		t = math.Abs(b) * epsilon
	} else {
		t = math.Abs(a) * epsilon
	}
	return math.Abs(a-b) <= t
}

func EssentiallyEqual(a, b, epsilon float64) bool {
	var t float64
	if math.Abs(a) > math.Abs(b) {
		t = math.Abs(b) * epsilon
	} else {
		t = math.Abs(a) * epsilon
	}
	return math.Abs(a-b) <= t
}

func DefinitelyGreaterThan(a, b, epsilon float64) bool {
	var t float64
	if math.Abs(a) < math.Abs(b) {
		t = math.Abs(b) * epsilon
	} else {
		t = math.Abs(a) * epsilon
	}
	return (a - b) > t
}

func DefinitelyLessThan(a, b, epsilon float64) bool {
	var t float64
	if math.Abs(a) < math.Abs(b) {
		t = math.Abs(b) * epsilon
	} else {
		t = math.Abs(a) * epsilon
	}
	return (a - b) < t
}
