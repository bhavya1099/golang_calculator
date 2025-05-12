package calc

import (
	math "math"
	testing "testing"

	assert "github.com/stretchr/testify/assert"
)

var testData = []multiplyTestData{
	{5.0, 3.0, 15.0},
	{5.0, -3.0, -15.0},
	{-5.0, -3.0, 15.0},
	{5.0, 0, 0},
}

type multiplyTestData struct {
	num1, num2, expectedResult float64
}

func TestDivideGoTest(t *testing.T) {

	var got float64

	got = Divide(6, 3)
	if math.Abs(got-2) > 1e-6 {
		t.Errorf("Divide(6, 3) = %v; want 2", got)
	}

	got = Divide(-6, -3)
	if math.Abs(got-2) > 1e-6 {
		t.Errorf("Divide(-6, -3) = %v; want 2", got)
	}

	got = Divide(-6, 3)
	if math.Abs(got+2) > 1e-6 {
		t.Errorf("Divide(-6, 3) = %v; want -2", got)
	}

	got = Divide(6, -3)
	if math.Abs(got+2) > 1e-6 {
		t.Errorf("Divide(6, -3) = %v; want -2", got)
	}

	got = Divide(6, 0)
	if !math.IsInf(got, 1) {
		t.Errorf("Divide(6, 0) = %v; want +Inf", got)
	}

}

func TestMultiply(t *testing.T) {
	for _, test := range testData {
		t.Run("Testing Multiply function", func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test: %v", r)
					t.Fail()
				}
			}()

			result := Multiply(test.num1, test.num2)

			assert.Equal(t, test.expectedResult, result)

			t.Logf("Passed. Expected: %f, Got: %f", test.expectedResult, result)
		})
	}
}

func TestSubtract(t *testing.T) {
	scenarios := []struct {
		desc   string
		num1   int
		num2   int
		result int
	}{
		{
			desc:   "Subtracting a positive number from a larger positive number",
			num1:   10,
			num2:   5,
			result: 5,
		},
		{
			desc:   "Subtracting a number from itself",
			num1:   8,
			num2:   8,
			result: 0,
		},
		{
			desc:   "Subtracting a positive number from a smaller positive number",
			num1:   5,
			num2:   10,
			result: -5,
		},
		{
			desc:   "Subtracting a number from zero",
			num1:   0,
			num2:   5,
			result: -5,
		},
		{
			desc:   "Subtracting zero from a number",
			num1:   10,
			num2:   0,
			result: 10,
		},
	}

	for _, s := range scenarios {
		t.Run(s.desc, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v", r)
					t.Fail()
				}
			}()

			result := Subtract(s.num1, s.num2)

			if result != s.result {
				t.Errorf("Expected: %v, got: %v", s.result, result)
			} else {
				t.Logf("Success: Expected %v and got %v", s.result, result)
			}
		})
	}
}
