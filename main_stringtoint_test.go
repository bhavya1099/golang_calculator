package golang_calculator

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"runtime/debug"
)

func TestStringToInt(t *testing.T) {
	t.Helper()
	// Use os.Stdout capture for observing function output
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
			t.Fail()
		}
	}()

	// Utility for capturing printed output
	tempStdout := os.Stdout
	r, w, _ := os.Pipe()
	defer func() { os.Stdout = tempStdout }()

	os.Stdout = w

	// Define test cases following table-driven tests
	tests := []struct {
		name       string
		input      string
		expOutput  int
		expError   bool
		validation func(resultOut int)
		outEvidence string
	}{
		{
			name:  "Test Normal String Conversion",
			input: "123",
			expOutput: 123,
			expError: false,
		}},
	}


test_)
}
