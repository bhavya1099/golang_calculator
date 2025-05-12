package calc

import (
	fmt "fmt"
	math "math"
	testing "testing"
)

func TestDivide(t *testing.T) {
	var tests = []struct {
		num1, num2, want float64
	}{
		{10, 5, 2.0},
		{-10, -5, 2.0},
		{-10, 5, -2.0},
		{10, 0, math.Inf(1)},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.num1, tt.num2)
		t.Run(testname, func(t *testing.T) {
			ans := Divide(tt.num1, tt.num2)
			if tt.num2 == 0 && (ans == math.Inf(1) || ans == math.Inf(-1)) {
				return
			} else if math.Abs(ans-tt.want) > 1e-6 {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
