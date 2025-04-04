package calc

import (
	testing "testing"
	math "math"
	assert "github.com/stretchr/testify/assert"
)



var testData = []multiplyTestData{
	{5.0, 3.0, 15.0},
	{5.0, -3.0, -15.0},
	{-5.0, -3.0, 15.0},
	{5.0, 0, 0},
}

type AddTestCase struct {
	Num1     int
	Num2     int
	Expected int
	Description string
}
type multiplyTestData struct {
	num1, num2, expectedResult float64
}


/*
ROOST_METHOD_HASH=Add_38f6779755
ROOST_METHOD_SIG_HASH=Add_8e349a90e1

FUNCTION_DEF=func Add(num1, num2 int) int 

*/
func TestAdd(t *testing.T) {
	testCases := []AddTestCase{
		{
			Num1:        2,
			Num2:        3,
			Expected:    5,
			Description: "Testing Positive Number Addition",
		},
		{
			Num1:        -2,
			Num2:        -3,
			Expected:    -5,
			Description: "Testing Negative Number Addition",
		},
		{
			Num1:        0,
			Num2:        5,
			Expected:    5,
			Description: "Adding Zero",
		},
		{
			Num1:        999999999,
			Num2:        999999999,
			Expected:    1999999998,
			Description: "Adding Large Numbers",
		},
		{
			Num1:        5,
			Num2:        -3,
			Expected:    2,
			Description: "Mixed Addition (Positive and Negative)",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test : %v", r)
					t.Fail()
				}
			}()

			result := Add(testCase.Num1, testCase.Num2)

			if result != testCase.Expected {
				t.Errorf("Failed: %s:\n For inputs num1: %d, num2: %d\n expected: %d, \n got: %d", testCase.Description, testCase.Num1, testCase.Num2, testCase.Expected, result)
			} else {
				t.Logf("Success: %s:\n For inputs num1: %d, num2: %d\n expected: %d, \n got: %d", testCase.Description, testCase.Num1, testCase.Num2, testCase.Expected, result)
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
	type args struct {
		num1 float64
		num2 float64
	}
	tests := []struct {
		name        string
		args        args
		want        float64
		description string
	}{
		{
			name: "Test case 1: Divide Positive Numbers",
			args: args{
				num1: 10.0,
				num2: 2.0,
			},
			want:        5.0,
			description: "This test will validate if the function can accurately divide two positive numbers.",
		},
		{
			name: "Test case 2: Divide by zero",
			args: args{
				num1: 10.0,
				num2: 0.0,
			},
			want:        math.Inf(1),
			description: "This test checks for error handling when dividing by zero. The result should be +Inf.",
		},
		{
			name: "Test case 3: Divide Zero by Positive Number",
			args: args{
				num1: 0.0,
				num2: 10.0,
			},
			want:        0.0,
			description: "This test checks the result when 0 is divided by a positive number. The result should be 0.",
		},
		{
			name: "Test case 4: Divide Two Negative Numbers",
			args: args{
				num1: -10.0,
				num2: -2.0,
			},
			want:        5.0,
			description: "This test checks whether the function can accurately divide two negative numbers. Dividing two negative numbers should result in a positive number.",
		},
		{
			name: "Test case 5: Divide A Positive and A Negative Number",
			args: args{
				num1: 10.0,
				num2: -2.0,
			},
			want:        -5.0,
			description: "This test checks whether the function can accurately divide a positive number by a negative number. The result should be negative.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered during test execution: %v", r)
					t.Fail()
				}
			}()
			if got := Divide(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			} else {
				t.Logf("Success: %s", tt.description)
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


/*
ROOST_METHOD_HASH=Subtract_559013d27f
ROOST_METHOD_SIG_HASH=Subtract_29b74c09c9

FUNCTION_DEF=func Subtract(num1, num2 int) int 

*/
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

