package tests

import (
	"scripts"
	"testing"
)

type calcTest struct {
	arg1, arg2, expected int
}

var addTests = []calcTest{
	calcTest{1, 1, 2},
	calcTest{2, 2, 4},
	calcTest{3, 3, 6},
}

var subTests = []calcTest{
	calcTest{2, 1, 1},
	calcTest{3, 1, 2},
	calcTest{4, 1, 3},
}

func TestCalc(t *testing.T) {

	for _, test := range addTests {
		if output := scripts.Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}

	for _, test := range subTests {
		if output := scripts.Subtract(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
