package main

import (
	fmt "fmt"
	os "os"
	strconv "strconv"
	testing "testing"
	debug "runtime/debug"
)








/*
ROOST_METHOD_HASH=stringToFloat64_d38659cd50
ROOST_METHOD_SIG_HASH=stringToFloat64_44e80853e6

FUNCTION_DEF=func stringToFloat64(str string) float64 

*/
func TestStringToFloat64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected interface{}
	}{
		{name: "Convert valid numeric string", input: "123.45", expected: float64(123.45)},
		{name: "Convert negative numeric string", input: "-123.45", expected: float64(-123.45)},
		{name: "Convert zero string", input: "0", expected: float64(0.0)},
		{name: "Empty string input", input: "", expected: 2},
		{name: "Scientific notation string", input: "1.23e2", expected: float64(123.0)},
		{name: "Invalid numeric content", input: "abc", expected: 2},
		{name: "String with whitespace", input: " 123.45 ", expected: float64(123.45)},
		{name: "Extremely large numeric string", input: "1e308", expected: float64(1e308)},
		{name: "Overflow numeric string", input: "1e309", expected: 2},
		{name: "Mixed valid and invalid characters", input: "123abc", expected: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered for test '%s'. Recovering gracefully.\nReason: %v\n%s", test.name, r, debug.Stack())
					t.Fail()
				}
			}()

			stdout := &os.File{}
			oldStdOut := os.Stdout
			defer func() {
				os.Stdout = oldStdOut
				t.Log("Restored os.Stdout successfully")
			}()
			os.Stdout = stdout

			var result float64
			var exitCode int
			func() {
				defer func() {
					if r := recover(); r != nil {
						exitCode = 2
					}
				}()

				result = stringToFloat64(test.input)
			}()

			if exitCode == 2 {
				if test.expected != 2 {
					t.Errorf("Test '%s' failed: expected no error, but got error (os.Exit code %d)", test.name, exitCode)
				} else {
					t.Logf("Test '%s' passed: os.Exit code matched expected (%d)", test.name, exitCode)
				}
			} else if expectedFloat, ok := test.expected.(float64); ok {
				if result != expectedFloat {
					t.Errorf("Test '%s' failed: expected %f, got %f", test.name, expectedFloat, result)
				} else {
					t.Logf("Test '%s' passed: expected value matched (%f)", test.name, expectedFloat)
				}
			} else {
				t.Errorf("Test '%s' invalid test case definition", test.name)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=stringToInt_73b9cbccee
ROOST_METHOD_SIG_HASH=stringToInt_e7cc66ec50

FUNCTION_DEF=func stringToInt(str string) int 

*/
func TestStringToInt(t *testing.T) {

	testCases := []struct {
		name     string
		input    string
		expected int
		wantErr  bool
	}{
		{
			name:     "Valid positive integer string",
			input:    "123",
			expected: 123,
			wantErr:  false,
		},
		{
			name:     "Valid zero string",
			input:    "0",
			expected: 0,
			wantErr:  false,
		},
		{
			name:     "Invalid string input",
			input:    "abc",
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "Negative integer string",
			input:    "-123",
			expected: -123,
			wantErr:  false,
		},
		{
			name:     "Empty input string",
			input:    "",
			expected: 0,
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered, test failed with input '%v'. Reason: %v\n%s", tc.input, r, string(debug.Stack()))
					t.Fail()
				}
			}()

			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("Failed to create pipe for stdout redirection: %v", err)
			}
			os.Stdout = w

			var result int
			func() {
				defer func() {
					if r := recover(); r != nil {

						t.Logf("os.Exit called for input '%s', expected error scenario", tc.input)
					}
				}()

				result = stringToInt(tc.input)
			}()

			w.Close()
			os.Stdout = os.NewFile(uintptr(r.Fd()), "/dev/tty")

			if tc.wantErr {

				if result != tc.expected {
					t.Errorf("Expected os.Exit on invalid input '%v', but function returned '%v'", tc.input, result)
				}
			} else {

				if result != tc.expected {
					t.Errorf("Expected %d for input '%v', but got %d", tc.expected, tc.input, result)
				} else {
					t.Logf("Success: Input '%v' converted to integer %d correctly.", tc.input, result)
				}
			}
		})
	}
}

