package calculator

import (
	"testing"
	"fmt"
)

func TestCalculator_Sum(t *testing.T) {

	calc := Calculator{}

	sum := calc.Sum(1, 3, 7)

	if (sum != 11.0) {
		t.Errorf("Expected %0.2f, but got %.2f", 11.0, sum)
	}

	sum = calc.Sum(3)
	if (sum != 3.0) {
		t.Errorf("Expected %0.2f, but got %.2f", 11.0, sum)
	}
}



func TestCalculator_Average(t *testing.T) {

	type testPair struct {
		values []float64
		answer float64
		name string
	}

	testCases := []testPair{
		{answer: 4.0, values:[]float64{2, 4, 6}, name: "first"},
		{answer: 0, values:[]float64{0}, name: "second"},
	}

	calc := Calculator{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if ans := calc.Average(tc.values...); ans != tc.answer {
				t.Errorf("expected %0.2f, got %0.2f", tc.answer, ans)
			}
		})
	}
}

func BenchmarkSprint(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = fmt.Sprint("Hello World")
	}
}

func BenchmarkSprintf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = fmt.Sprintf("Hello World")
	}
}

func TestCalculator_Multiply(t *testing.T) {
	type testPair struct {
		values []float64
		answer float64
	}

	testCases := []testPair {
		{answer:21.0, values:[]float64{1,3,7}},
		{answer:0.0, values:[]float64{}},
	}

	calc := Calculator{}
	for _,tc := range testCases {
		ans := calc.Multiply(tc.values...);
		if ans != tc.answer {
			t.Errorf("Expected %0.2f, but got %.2f", tc.answer, ans)
		}
		//if ans := calc.Multiply(tc.values...); ans != tc.answer {
		//	t.Errorf("Expected %0.2f, but got %.2f", tc.answer, ans)
		//}

	}
}
