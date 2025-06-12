package calc

import (
	"fmt"
	"math"
	"os"
	"testing"
)

func TestMultiply(t *testing.T) {
	tests := []struct {
		name         string
		num1         float64
		num2         float64
		expected     float64
		shouldPanic  bool // Flag to check panic scenarios
	}{
		// תרחים 1: בדיקת פעולת הכפל עם מספרים רגילים חיוביים
		{
			name:     "Multiply two positive numbers",
			num1:     5.0,
			num2:     2.0,
			expected: 10.0,
		},
		// תרחים 2: בדיקת פעולת הכפל כאשר אחד מהמספרים הוא אפס
		{
			name:     "Multiply with zero",
			num1:     0.0,
			num2:     5.0,
			expected: 0.0,
		},
		// תרחים 3: בדיקת פעולת הכפל כאשר שני המספרים שליליים
		{
			name:     "Multiply two negative numbers",
			num1:     -3.0,
			num2:     -2.0,
			expected: 6.0,
		},
		// תרחים 4: בדיקת פעולת הכפל כאשר אחד המספרים שלילי והשני חיובי
		{
			name:     "Multiply negative and positive number",
			num1:     -3.0,
			num2:     4.0,
			expected: -12.0,
		},
		// תרחים 5: בדיקת פעולת הכפל כאשר אחד המספרים קרוב ל-0 ואחד גדול מאוד
		{
			name:     "Multiply number close to 0 with a large number",
			num1:     0.0000001,
			num2:     1000000.0,
			expected: 0.1,
		},
		// תרחים 6: בדיקה עבור ערכים לא תקניים כמו אינסוף
		{
			name:     "Multiply with infinity",
			num1:     math.Inf(1),
			num2:     5.0,
			expected: math.Inf(1),
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

			// Capture os.Stdout for testing functions that write to output
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Testing Multiply function
			result := Multiply(tt.num1, tt.num2)

			// Restore original os.Stdout
			w.Close()
			os.Stdout = oldStdout
			var output string
			fmt.Fscanf(r, "%s", &output)

			if tt.shouldPanic {
				t.Errorf("Expected panic but did not encounter one")
			} else {
				// Log detailed output for debugging
				if result != tt.expected {
					t.Errorf("Test failed: %s\nExpected: %v\nGot: %v", tt.name, tt.expected, result)
				} else {
					t.Logf("Test passed: %s\nExpected and Got: %v", tt.name, result)
				}
			}
		})
	}
}
