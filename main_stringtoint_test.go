package golang_calculator

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"testing"
	"runtime/debug"
)

func TestStringToInt(t *testing.T) {
	// Capture os.Stdout for reliable output testing
	var stdoutBuffer bytes.Buffer
	oldStdout := os.Stdout
	os.Stdout = &stdoutBuffer // Redirect standard output to memory buffer
	defer func() { os.Stdout = oldStdout }() // Restore original stdout at the end of the test

	// Table-driven tests
	testCases := []struct {
		name          string
		input         string
		expected      int
		expectPanic   bool
		expectedLog   string
	}{
		{
			name:        "Valid String Integer",
			input:       "123",
			expected:    123,
			expectPanic: false,
			expectedLog: "Function should return 123 for valid input '123'.",
		},
		{
			name:        "Empty String",
			input:       "",
			expected:    0, // Returning 0 is irrelevant since a panic will occur
			expectPanic: true,
			expectedLog: "Empty input string should cause program exit.",
		},
		{
			name:        "String with non-numeric characters",
			input:       "abc",
			expected:    0, // Returning 0 is irrelevant since a panic will occur
			expectPanic: true,
			expectedLog: "Input 'abc' should cause program exit due to non-numeric data.",
		},
		{
			name:        "Negative Integer String",
			input:       "-123",
			expected:    -123,
			expectPanic: false,
			expectedLog: "Negative input '-123' should return -123.",
		},
		{
			name:        "Number with leading zeroes",
			input:       "000123",
			expected:    123,
			expectPanic: false,
			expectedLog: "Input '000123' should be converted to 123 correctly ignoring leading zeroes.",
		},
		{
			name:        "Overflow number",
			input:       "999999999999999999999999",
			expected:    0, // Returning 0 is irrelevant since a panic will occur
			expectPanic: true,
			expectedLog: "Large input '999999999999999999999999' should cause program exit due to overflow.",
		},
		{
			name:        "Input with leading and trailing spaces",
			input:       "   123   ",
			expected:    123,
			expectPanic: false,
			expectedLog: "Input '   123   ' should be parsed into 123 ignoring surrounding spaces.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Defer to recover from potential panics
			defer func() {
				if r := recover(); r != nil {
					if tc.expectPanic {
						t.Logf("[Scenario '%s'] Panic as expected: %v\n%s", tc.name, r, string(debug.Stack()))
						t.Log(tc.expectedLog)
					} else {
						t.Errorf("[Scenario '%s'] Unexpected panic occurred: %v\n%s", tc.name, r, string(debug.Stack()))
					}
				}
			}()

			// Clean previous stdout buffer data
			stdoutBuffer.Reset()

			// Actual test for the scenario
			result := stringToInt(tc.input)

			if !tc.expectPanic {
				if result != tc.expected {
					t.Errorf("Test failed. Expected: %d, Got: %d. %s", tc.expected, result, tc.expectedLog)
				} else {
					t.Logf("Test passed. Expected: %d, Got: %d. %s", tc.expected, result, tc.expectedLog)
				}
			} else {
				// The buffer will contain any output by the panic-producing cases
				if !bytes.Contains(stdoutBuffer.Bytes(), []byte(tc.input)) {
					t.Errorf("Expected program to print error containing '%s', but found: %s", tc.input, stdoutBuffer.String())
				}
			}
		})
	}
}

// TODO: Generalize the test patterns to work with additional variants or edge cases as needed in the future.
