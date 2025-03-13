package calc

import (
	math "math"
	testing "testing"
	os "os"
	debug "runtime/debug"
)

/*
ROOST_METHOD_HASH=Add_38f6779755
ROOST_METHOD_SIG_HASH=Add_8e349a90e1

FUNCTION_DEF=func Add(num1, num2 int) int

*/
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		num1     int
		num2     int
		expected int
	}{
		{
			name:     "Adding Two Positive Integers",
			num1:     5,
			num2:     10,
			expected: 15,
		},
		{
			name:     "Adding Two Negative Integers",
			num1:     -3,
			num2:     -7,
			expected: -10,
		},
		{
			name:     "Adding a Positive Integer and a Negative Integer",
			num1:     8,
			num2:     -3,
			expected: 5,
		},
		{
			name:     "Adding Zero to an Integer",
			num1:     0,
			num2:     12,
			expected: 12,
		},
		{
			name:     "Adding Two Zeros",
			num1:     0,
			num2:     0,
			expected: 0,
		},
		{
			name:     "Adding Max Int and 1",
			num1:     math.MaxInt32,
			num2:     1,
			expected: math.MinInt32,
		},
		{
			name:     "Adding Min Int and -1",
			num1:     math.MinInt32,
			num2:     -1,
			expected: math.MaxInt32,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			old := os.Stdout
			_, w, _ := os.Pipe()
			os.Stdout = w

			result := Add(tt.num1, tt.num2)

			w.Close()
			os.Stdout = old

			if result != tt.expected {
				t.Errorf("Test failed for %s: expected %d, got %d", tt.name, tt.expected, result)
			} else {
				t.Logf("Test succeeded for %s: expected and got %d", tt.name, tt.expected)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=Divide_052b9c25ea
ROOST_METHOD_SIG_HASH=Divide_15b7594322

FUNCTION_DEF=func Divide(num1, num2 float64) float64

*/
func TestDivide(t *testing.T) {

	oldStdout := os.Stdout
	_, w, _ := os.Pipe()
	os.Stdout = w

	tests := []struct {
		name     string
		num1     float64
		num2     float64
		expected float64
	}{
		{
			name:     "Division of Two Positive Numbers",
			num1:     10.0,
			num2:     2.0,
			expected: 5.0,
		},
		{
			name:     "Division Resulting in a Fraction",
			num1:     7.0,
			num2:     2.0,
			expected: 3.5,
		},
		{
			name:     "Division by Zero",
			num1:     5.0,
			num2:     0.0,
			expected: 0.0,
		},
		{
			name:     "Division of Zero by a Positive Number",
			num1:     0.0,
			num2:     5.0,
			expected: 0.0,
		},
		{
			name:     "Division of Two Negative Numbers",
			num1:     -10.0,
			num2:     -2.0,
			expected: 5.0,
		},
		{
			name:     "Division of a Positive Number by a Negative Number",
			num1:     10.0,
			num2:     -2.0,
			expected: -5.0,
		},
		{
			name:     "Division of a Negative Number by a Positive Number",
			num1:     -10.0,
			num2:     2.0,
			expected: -5.0,
		},
		{
			name:     "Division Resulting in a Very Small Fraction",
			num1:     1.0,
			num2:     10000.0,
			expected: 0.0001,
		},
		{
			name:     "Division of a Large Number by a Small Number",
			num1:     1000000.0,
			num2:     0.0001,
			expected: 10000000000.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			result := Divide(tt.num1, tt.num2)

			if (tt.num2 == 0.0 && result != tt.expected) || (tt.num2 != 0.0 && result != tt.expected) {
				t.Errorf("Test failed: %s. Expected: %v, Got: %v", tt.name, tt.expected, result)
			} else {
				t.Logf("Test passed: %s. Expected: %v, Got: %v", tt.name, tt.expected, result)
			}
		})
	}

	w.Close()
	os.Stdout = oldStdout
}

/*
ROOST_METHOD_HASH=Multiply_1585632006
ROOST_METHOD_SIG_HASH=Multiply_d6ab1fb07f

FUNCTION_DEF=func Multiply(num1, num2 float64) float64

*/
func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		num1     float64
		num2     float64
		expected float64
	}{
		{
			name:     "Multiply two positive numbers",
			num1:     3.5,
			num2:     2.0,
			expected: 7.0,
		},
		{
			name:     "Multiply a positive and a negative number",
			num1:     3.5,
			num2:     -2.0,
			expected: -7.0,
		},
		{
			name:     "Multiply two negative numbers",
			num1:     -3.5,
			num2:     -2.0,
			expected: 7.0,
		},
		{
			name:     "Multiply a number by zero",
			num1:     3.5,
			num2:     0.0,
			expected: 0.0,
		},
		{
			name:     "Multiply a number by one",
			num1:     3.5,
			num2:     1.0,
			expected: 3.5,
		},
		{
			name:     "Multiply two very large numbers",
			num1:     1e308,
			num2:     2.0,
			expected: math.Inf(1),
		},
		{
			name:     "Multiply two very small numbers",
			num1:     1e-308,
			num2:     2.0,
			expected: 2e-308,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			result := Multiply(tt.num1, tt.num2)
			if result != tt.expected {
				t.Errorf("Test %s failed: expected %v, got %v", tt.name, tt.expected, result)
			} else {
				t.Logf("Test %s succeeded: expected %v, got %v", tt.name, tt.expected, result)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=Subtract_559013d27f
ROOST_METHOD_SIG_HASH=Subtract_29b74c09c9

FUNCTION_DEF=func Subtract(num1, num2 int) int

*/
func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		num1     int
		num2     int
		expected int
	}{
		{"Subtracting two positive integers", 10, 5, 5},
		{"Subtracting a larger number from a smaller number", 5, 10, -5},
		{"Subtracting zero from a number", 7, 0, 7},
		{"Subtracting a number from itself", 8, 8, 0},
		{"Subtracting negative numbers", -3, -7, 4},
		{"Subtracting a negative number from a positive number", 10, -5, 15},
		{"Subtracting a positive number from a negative number", -8, 3, -11},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			result := Subtract(tt.num1, tt.num2)
			if result != tt.expected {
				t.Errorf("Subtract(%d, %d) = %d; want %d", tt.num1, tt.num2, result, tt.expected)
			} else {
				t.Logf("Success: Subtract(%d, %d) = %d", tt.num1, tt.num2, result)
			}
		})
	}
}
