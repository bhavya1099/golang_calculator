package calc

import (
	fmt "fmt"
	math "math"
	os "os"
	debug "runtime/debug"
	sync "sync"
	testing "testing"
	strings "strings"
)







func TestAdd(t *testing.T) {
	type testCase struct {
		name     string
		input1   int
		input2   int
		expected int
	}

	tests := []testCase{
		{
			name:     "Adding two positive integers",
			input1:   5,
			input2:   10,
			expected: 15,
		},
		{
			name:     "Adding two negative integers",
			input1:   -3,
			input2:   -7,
			expected: -10,
		},
		{
			name:     "Adding a positive integer and a negative integer",
			input1:   7,
			input2:   -2,
			expected: 5,
		},
		{
			name:     "Adding zero to a number",
			input1:   8,
			input2:   0,
			expected: 8,
		},
		{
			name:     "Adding two zero values",
			input1:   0,
			input2:   0,
			expected: 0,
		},
		{
			name:     "Adding integers with the same magnitude but opposite signs",
			input1:   15,
			input2:   -15,
			expected: 0,
		},
		{
			name:     "Adding integers in sequence (Associative Property)",
			input1:   3,
			input2:   5,
			expected: 8,
		},

		{
			name:     "Adding integers that result in boundary values",
			input1:   math.MaxInt32,
			input2:   -1,
			expected: math.MaxInt32 - 1,
		},
		{
			name:     "Large integer addition with potential overflow",
			input1:   math.MaxInt32,
			input2:   1,
			expected: math.MinInt32,
		},
		{
			name:     "Adding minimum integer values",
			input1:   math.MinInt32,
			input2:   -1,
			expected: math.MinInt32 - 1,
		},
	}

	var wg sync.WaitGroup
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered: %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			wg.Add(1)
			go func() {
				defer wg.Done()

				result := Add(tc.input1, tc.input2)
				if result != tc.expected {
					t.Errorf("Test %s FAILED: expected %d, got %d", tc.name, tc.expected, result)
				} else {
					t.Logf("Test %s PASSED: expected %d, got %d", tc.name, tc.expected, result)
				}

				if tc.name == "Adding integers in sequence (Associative Property)" {
					groupedResult1 := Add(Add(tc.input1, tc.input2), 7)
					groupedResult2 := Add(tc.input1, Add(tc.input2, 7))
					if groupedResult1 != groupedResult2 {
						t.Errorf("Associative property FAILED for %s: group1=%d, group2=%d",
							tc.name, groupedResult1, groupedResult2)
					} else {
						t.Logf("Associative property PASSED for %s: group1=%d, group2=%d",
							tc.name, groupedResult1, groupedResult2)
					}
				}
			}()
		})
	}
	wg.Wait()

	stdoutBackup := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fmt.Fprintf(w, "Simulating output from function...\n")
	w.Close()
	os.Stdout = stdoutBackup

	var output string
	fmt.Fscanf(r, "%s\n", &output)

	if output != "Simulating output from function..." {
		t.Errorf("Unexpected output in stdout test: got %s", output)
	}
	r.Close()
}

func TestDivide(t *testing.T) {
	type testCase struct {
		num1   float64
		num2   float64
		expect float64
		name   string
		errMsg string
	}

	testCases := []testCase{
		{
			num1:   10,
			num2:   2,
			expect: 5,
			name:   "Scenario 1: Division of two positive numbers",
			errMsg: "Expected num1 / num2 to match the given result.",
		},
		{
			num1:   10,
			num2:   -2,
			expect: -5,
			name:   "Scenario 2: Division of a positive number by a negative number",
			errMsg: "Expected negative quotient due to opposite signs.",
		},
		{
			num1:   -10,
			num2:   2,
			expect: -5,
			name:   "Scenario 3: Division of a negative number by a positive number",
			errMsg: "Expected negative quotient for mixed signs.",
		},
		{
			num1:   -10,
			num2:   -2,
			expect: 5,
			name:   "Scenario 4: Division of two negative numbers",
			errMsg: "Expected positive quotient because both signs are negative.",
		},
		{
			num1:   0,
			num2:   2,
			expect: 0,
			name:   "Scenario 6: Division of zero by a non-zero number",
			errMsg: "Expected quotient to be zero.",
		},
		{
			num1:   1e308,
			num2:   1e-308,
			expect: math.MaxFloat64,
			name:   "Scenario 7: Division of a very large number by a very small number",
			errMsg: "Expected quotient to approach large numbers.",
		},
		{
			num1:   1e-308,
			num2:   1e308,
			expect: 0,
			name:   "Scenario 8: Division of a very small number by a very large number",
			errMsg: "Expected quotient to approach zero.",
		},
		{
			num1:   math.MaxFloat64,
			num2:   math.SmallestNonzeroFloat64,
			expect: math.Inf(1),
			name:   "Scenario 9: Division of extremely small and extremely large float values",
			errMsg: "Expected the result to properly handle overflow.",
		},
		{
			num1:   10,
			num2:   10,
			expect: 1,
			name:   "Scenario 10: Division of identical non-zero numbers",
			errMsg: "Expected quotient to be exactly one for identical inputs.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered: %v", r)
					t.Fail()
				}
			}()

			got := Divide(tc.num1, tc.num2)
			if math.IsNaN(got) || math.IsInf(got, 0) {
				t.Logf("Division result produced NaN/Inf requiring specific handling. Inputs: num1=%f, num2=%f", tc.num1, tc.num2)
			}

			if math.Abs(got-tc.expect) > 1e-10 {
				t.Errorf("FAILED %s: Got '%f', expected '%f'. %s", tc.name, got, tc.expect, tc.errMsg)
			} else {
				t.Logf("PASSED %s: Got consistent result '%f'", tc.name, got)
			}
		})
	}

	t.Run("Scenario 5: Division by zero", func(t *testing.T) {
		defer func() {
			if recovery := recover(); recovery != nil {
				t.Logf("Panic from division by zero caught: %v", recovery)
			}
		}()

		got := Divide(100, 0)
		if !math.IsInf(got, 1) && !(math.IsNaN(got)) {
			t.Errorf("FAILED Division by zero management: Unexpected output '%f'", got)
		} else {
			t.Logf("PASSED Division by zero: Handled correctly.")
		}
	})
}

func TestSubtract(t *testing.T) {

	tests := []struct {
		name     string
		num1     int
		num2     int
		expected int
		desc     string
	}{
		{
			name:     "Subtracting two positive integers",
			num1:     10,
			num2:     5,
			expected: 5,
			desc:     "Subtracting smaller positive integers verifies basic functionality for valid input ranges.",
		},
		{
			name:     "Subtracting two negative integers",
			num1:     -10,
			num2:     -5,
			expected: -5,
			desc:     "Negative integer arithmetic validates subtraction operation sign conventions.",
		},
		{
			name:     "Subtracting zero from a positive integer",
			num1:     15,
			num2:     0,
			expected: 15,
			desc:     "Identity property of subtraction should return the other operand unchanged.",
		},
		{
			name:     "Subtracting zero from a negative integer",
			num1:     -20,
			num2:     0,
			expected: -20,
			desc:     "The number remains unchanged when zero is subtracted.",
		},
		{
			name:     "Subtracting a positive integer from zero",
			num1:     0,
			num2:     7,
			expected: -7,
			desc:     "Zero minus a positive number should correctly yield a negative number.",
		},
		{
			name:     "Subtracting a negative integer from zero",
			num1:     0,
			num2:     -4,
			expected: 4,
			desc:     "Handling a negative subtraction from zero yields positive counterpart.",
		},
		{
			name:     "Subtracting two large positive integers",
			num1:     1000000,
			num2:     999999,
			expected: 1,
			desc:     "Testing large integer subtraction ensures capability in high-value ranges.",
		},
		{
			name:     "Subtracting two large negative integers",
			num1:     -1000000,
			num2:     -999999,
			expected: -1,
			desc:     "Handling large-value extreme negative bounds verifies proper implementation.",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			t.Logf("Starting test: %s | Description: %s", test.name, test.desc)

			actual := Subtract(test.num1, test.num2)

			if actual != test.expected {
				t.Errorf("FAILED: %s | Got: %d, Expected: %d", test.name, actual, test.expected)
			} else {
				t.Logf("PASSED: %s | Got: %d, Expected: %d", test.name, actual, test.expected)
			}
		})
	}

	var stdoutBuf strings.Builder

	t.Log("TestSubtract execution completed.")
	fmt.Fprintf(&stdoutBuf, "TestSubtract execution completed.\n")

	t.Log(stdoutBuf.String())
}

