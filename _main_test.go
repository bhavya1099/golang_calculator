
// ********RoostGPT********
/*

roost_feedback [02/07/2025, 3:38:01 PM]:remove compilation errors if any\n\n
*/

// ********RoostGPT********

package main

import (
	"os"
	"strconv"
	"testing"
	"runtime/debug"
)

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
					t.Errorf("Panic encountered for test '%s'. Reason: %v\n%s", test.name, r, debug.Stack())
				}
			}()

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
				}
			} else if expectedFloat, ok := test.expected.(float64); ok {
				if result != expectedFloat {
					t.Errorf("Test '%s' failed: expected %f, got %f", test.name, expectedFloat, result)
				}
			}
		})
	}
}

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
					t.Errorf("Panic encountered, test failed with input '%v'. Reason: %v\n%s", tc.input, r, debug.Stack())
				}
			}()

			var result int
			func() {
				defer func() {
					if r := recover(); r != nil {
						t.Logf("os.Exit called for input '%s', expected error scenario", tc.input)
					}
				}()

				result = stringToInt(tc.input)
			}()

			if tc.wantErr {
				if result != tc.expected {
					t.Errorf("Expected os.Exit on invalid input '%v', but function returned '%v'", tc.input, result)
				}
			} else {
				if result != tc.expected {
					t.Errorf("Expected %d for input '%v', but got %d", tc.expected, tc.input, result)
				}
			}
		})
	}
}
