package internal

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func Test_scanData(t *testing.T) {
	tests := []struct {
		input              string
		expectedColors     []int
		expectedContainers []int
		expectedError      bool
	}{
		{
			input:              "3\n1 2 3\n1 2 3\n1 2 3\n",
			expectedColors:     []int{3, 6, 9},
			expectedContainers: []int{6, 6, 6},
			expectedError:      false,
		},
		{
			input:         "2\n1 2\n1 x\n1 2",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		in := bufio.NewReader(strings.NewReader(tt.input))
		colorBalls, containerBalls, err := scanData(in)
		if (err != nil) != tt.expectedError {
			t.Errorf("Get error = %v, expected error = %v", err, tt.expectedError)
			continue
		}
		if !tt.expectedError {
			if len(tt.expectedColors) != len(tt.expectedContainers) {
				t.Errorf("Invalid test data, wrong length")
				continue
			}
			if !isEqual(colorBalls, tt.expectedColors) {
				t.Errorf("Get colorBalls = %v, expected colorBalls = %v", colorBalls, tt.expectedColors)
			}
			if !isEqual(containerBalls, tt.expectedContainers) {
				t.Errorf("Get containerBalls = %v, expected containerBalls = %v", containerBalls, tt.expectedContainers)
			}
		}
	}
}

func Test_isEqual(t *testing.T) {
	tests := []struct {
		colorBalls     []int
		containerBalls []int
		expectedOutput bool
	}{
		{
			colorBalls:     []int{1, 2, 3},
			containerBalls: []int{1, 2, 4},
			expectedOutput: false,
		},
		{
			colorBalls:     []int{1, 2, 3},
			containerBalls: []int{1, 2, 3},
			expectedOutput: true,
		},
	}

	for _, tt := range tests {
		result := isEqual(tt.colorBalls, tt.containerBalls)
		if result != tt.expectedOutput {
			t.Errorf("Get = %v, expected output = %v", result, tt.expectedOutput)
		}
	}
}

func Test_printResult(t *testing.T) {
	tests := []struct {
		result         bool
		expectedOutput string
	}{
		{result: true, expectedOutput: "yes"},
		{result: false, expectedOutput: "no"},
	}

	for _, tt := range tests {
		var buf bytes.Buffer
		writer := bufio.NewWriter(&buf)
		printResult(writer, tt.result)
		writer.Flush()

		if buf.String() != tt.expectedOutput {
			t.Errorf("Get output = %v, expected output = %v", buf.String(), tt.expectedOutput)
		}
	}
}

func Test_sortBalls(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{
			input:          "2\n1 2\n2 1\n",
			expectedOutput: "yes",
		},
		{
			input:          "3\n10 20 30\n1 1 1\n0 0 1\n",
			expectedOutput: "no",
		},
		{
			input:          "3\n1 2 3\n4 5 6\n7 8 9\n",
			expectedOutput: "no",
		},

		{
			input:          "3\n1 1 1\n1 1 1\n1 1 1\n",
			expectedOutput: "yes",
		},
		{
			input:          "2\n1 0\n0 1\n",
			expectedOutput: "yes",
		},
		{
			input:          "2\n1 0\n0 0\n",
			expectedOutput: "yes",
		},
		{
			input:          "1\n1\n",
			expectedOutput: "yes",
		},
		{
			input:          "3\n0 0 0\n1 1 1\n 0 0 0",
			expectedOutput: "no",
		},
		{
			input:          "3\n1 0 0\n0 1 0\n 0 0 1",
			expectedOutput: "yes",
		},
	}

	for _, tt := range tests {
		in := bufio.NewReader(strings.NewReader(tt.input))
		var buf bytes.Buffer
		out := bufio.NewWriter(&buf)

		SortBalls(in, out)
		out.Flush()

		if buf.String() != tt.expectedOutput {
			t.Errorf("Get output = %v, expected output %v", buf.String(), tt.expectedOutput)
		}
	}
}
