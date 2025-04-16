package calc

import (
	fmt "fmt"
	math "math"
	os "os"
	debug "runtime/debug"
	strings "strings"
	testing "testing"

	assert "github.com/stretchr/testify/assert"
)

type TestData struct {
	Num1, Num2, ExpectedResult float64
}
type testCase struct {
	name       string
	num1       int
	num2       int
	expected   int
	shouldFail bool
}

/*
ROOST_METHOD_HASH=Add_38f6779755
ROOST_METHOD_SIG_HASH=Add_8e349a90e1

FUNCTION_DEF=func Add(num1, num2 int) int
*/
func TestAdd(t *testing.T) {

	testCases := []struct {
		name     string
		num1     int
		num2     int
		expected int
	}{
		{
			name:     "Testing addition of positive numbers",
			num1:     2,
			num2:     3,
			expected: 5,
		},
		{
			name:     "Testing addition of negative numbers",
			num1:     -2,
			num2:     -3,
			expected: -5,
		},
		{
			name:     "Testing addition of a positive and a negative number",
			num1:     2,
			num2:     -3,
			expected: -1,
		},
		{
			name:     "Testing addition leading to integer overflow",
			num1:     math.MaxInt64,
			num2:     math.MaxInt64,
			expected: -2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered failing the test: %v", r)
					t.Fail()
				}
			}()

			result := Add(tc.num1, tc.num2)

			if result != tc.expected {
				t.Errorf("Expected result: %v, Got: %v", tc.expected, result)
			} else {
				t.Logf("Success: Expected result: %v, Got: %v", tc.expected, result)
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
	type testData struct {
		num1              float64
		num2              float64
		expectedResult    float64
		expectInfinite    bool
		testDescription   string
		expectedToSucceed bool
	}

	testCases := []testData{
		{10.0, 5.0, 2.0, false, "Division of two positive numbers", true},
		{-10.0, -5.0, 2.0, false, "Division of two negative numbers", true},
		{10.0, 0.0, 0.0, true, "Division by zero", false},
	}

	for _, test := range testCases {
		t.Run(test.testDescription, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if test.expectedToSucceed {
						t.Errorf("The test case failed due to panic: %v", r)
					} else {
						t.Logf("A panic is expected in this test case.")
					}
				} else if !test.expectedToSucceed {
					t.Errorf("The test case was expected to panic but it didn't.")
				}
			}()

			result := Divide(test.num1, test.num2)

			if test.expectInfinite && !(math.IsInf(result, 1) || math.IsInf(result, -1)) {
				t.Errorf("Expected: Division by Zero to be +Inf or -Inf, but got: %v", result)
			} else if !test.expectInfinite && !assert.Equal(t, test.expectedResult, result) {
				t.Errorf("Expected: %v, but got: %v", test.expectedResult, result)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=Multiply_1585632006
ROOST_METHOD_SIG_HASH=Multiply_d6ab1fb07f

FUNCTION_DEF=func Multiply(num1, num2 float64) float64
*/
func TestMultiply(t *testing.T) {

	testCases := []struct {
		name     string
		data     TestData
		hasError bool
	}{
		{"Positive Numbers", TestData{5, 6, 30}, false},
		{"Negative Numbers", TestData{-5, -6, 30}, false},
		{"Positive and Negative Number", TestData{5, -6, -30}, false},
		{"Multiplication with Zero", TestData{5, 0, 0}, false},
		{"Floating Point Numbers", TestData{1.2, 3.6, 4.32}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Fatal(fmt.Sprintf("Test case [%s] panic occurred: %v \n%s", tc.name, r, string(debug.Stack())))
				}
			}()

			result := Multiply(tc.data.Num1, tc.data.Num2)
			if result != tc.data.ExpectedResult {
				t.Errorf("Test case [%s] failed. Expected result %v but got %v", tc.name, tc.data.ExpectedResult, result)
			} else {
				t.Logf("Test case [%s] passed.", tc.name)
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
	testCases := []testCase{
		{"Positive integers subtraction", 10, 5, 5, false},
		{"Negative integers subtraction", -8, -2, -6, false},
		{"Zero subtraction", 7, 0, 7, false},
		{"Integer subtracted from itself", 4, 4, 0, false},
		{"Large integer values subtraction", 1e6, 1e5, 9e5, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Panic occurred during execution: %v", r)
				}
			}()

			out := Subtract(tc.num1, tc.num2)

			if out != tc.expected && !tc.shouldFail {
				t.Errorf("Expected %v, but got %v", tc.expected, out)
			}

			if tc.shouldFail && out != tc.expected {
				t.Errorf("Expected test to fail, but it passed")
			}
		})
	}

	t.Run("Test writing to stdout", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Panic occurred during execution: %v", r)
			}
		}()

		reader, writer, err := os.Pipe()
		if err != nil {
			t.Fatal(err)
		}

		os.Stdout = writer

		outChan := make(chan string)
		go func() {
			var buf strings.Builder
			_, err := fmt.Fscanf(reader, "%s", &buf)
			if err != nil {
				t.Error(err)
			}
			outChan <- buf.String()
		}()

		num1 := 3
		num2 := 1
		Subtract(num1, num2)

		writer.Close()
		out := <-outChan

		expected := fmt.Sprintf("%d subtracted by %d is %d", num1, num2, num1-num2)
		if out == expected {
			t.Errorf("Expected %s, but got %s", expected, out)
		}
	})
}
