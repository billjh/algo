package misc

import (
	"fmt"
	"testing"
)

func TestGCDInvalidInput(t *testing.T) {
	testCases := []struct{ a, b uint }{
		{2, 0},
		{0, 2},
		{0, 0},
	}
	for _, testCase := range testCases {
		if GCD(testCase.a, testCase.b) != 0 {
			t.Error(fmt.Sprintf("GCD: undetected invalid input %d or %d", testCase.a, testCase.b))
		}
	}
}

func TestGCD(t *testing.T) {
	testCases := []struct{ a, b, div uint }{
		{12, 6, 6},
		{6, 12, 6},
		{4851, 3003, 231},
		{100, 7, 1},
		{3, 2, 1},
	}
	for _, testCase := range testCases {
		if GCD(testCase.a, testCase.b) != testCase.div {
			t.Error(fmt.Sprintf("GCD fails on %d, %d => expect %d get %d ", testCase.a, testCase.b, testCase.div, GCD(testCase.a, testCase.b)))
		}
	}
}
