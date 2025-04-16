package golang_calculator

import (
	fmt "fmt"
	ioutil "io/ioutil"
	os "os"
	strings "strings"
	testing "testing"
)








/*
ROOST_METHOD_HASH=stringToFloat64_d38659cd50
ROOST_METHOD_SIG_HASH=stringToFloat64_44e80853e6

FUNCTION_DEF=func stringToFloat64(str string) float64 

*/
func TestStringToFloat64(t *testing.T) {

	testCases := []struct {
		name        string
		input       string
		expected    float64
		expectError bool
	}{
		{name: "Valid numeric string", input: "123.456", expected: 123.456, expectError: false},
		{name: "Invalid numeric string", input: "abc", expectError: true},
		{name: "Empty string", input: "", expected: 0, expectError: false},
		{name: "Float64 string with excess precision", input: "123.456789123456789", expected: 123.456789123456789, expectError: false},
		{name: "Large magnitude float64 string", input: "1.8e308", expected: 1.8e308, expectError: false},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v", r)
					t.Fail()
				}
			}()

			r, w, _ := os.Pipe()
			os.Stdout = w

			actualResult := stringToFloat64(testCase.input)

			w.Close()
			out, _ := ioutil.ReadAll(r)
			os.Stdout = osOutBackup

			expectedError := fmt.Sprintf("strconv.ParseFloat: parsing \"%s\": invalid syntax", testCase.input)

			if testCase.expectError {

				if !strings.Contains(string(out), expectedError) {
					t.Errorf("Test '%s' failed. Expected error message: '%s', but got: '%s'", testCase.name, expectedError, string(out))
					return
				}
				t.Logf("Test '%s' passed successfully.", testCase.name)
				return

			}

			if actualResult != testCase.expected {
				t.Errorf("Test '%s' failed. Expected: '%f', but got:  '%f'", testCase.name, testCase.expected, actualResult)
				return
			}
			t.Logf("Test '%s' passed successfully.", testCase.name)
		})
	}
}


/*
ROOST_METHOD_HASH=stringToInt_73b9cbccee
ROOST_METHOD_SIG_HASH=stringToInt_e7cc66ec50

FUNCTION_DEF=func stringToInt(str string) int 

*/
func TestStringToInt(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput int
		expectError    bool
	}{
		{
			name:           "Scenario 1: Testing with a positive integer string",
			input:          "122",
			expectedOutput: 122,
			expectError:    false,
		},
		{
			name:           "Scenario 2: Testing with a negative integer string",
			input:          "-122",
			expectedOutput: -122,
			expectError:    false,
		},
		{
			name:           "Scenario 3: Testing with zero string",
			input:          "0",
			expectedOutput: 0,
			expectError:    false,
		},
		{
			name:        "Scenario 4: Testing with a non-integer string",
			input:       "abc",
			expectError: true,
		},
		{
			name:        "Scenario 5: Testing with a string containing integer and non-integer characters",
			input:       "123abc",
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n", r)
					t.Fail()
				}
			}()

			rescueStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			output := stringToInt(test.input)
			w.Close()

			out, _ := ioutil.ReadAll(r)
			os.Stdout = rescueStdout
			strOut := strings.TrimSpace(string(out))

			if test.expectError {
				if strOut == "" {
					t.Errorf("expected an error but got none")
				}
			} else if output != test.expectedOutput {
				t.Errorf("expected %d, but got %d", test.expectedOutput, output)
			}

		})
	}
}

