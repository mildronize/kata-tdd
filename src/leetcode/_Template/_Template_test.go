package _Template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name     string
	input    string
	expected bool
}{
	{"Name of test Case", "input_string", true},
}

func Test_Template(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(subtest *testing.T) {
			// subtest.Parallel()
			actual := _Template(test.input)
			assert.Equal(subtest, test.expected, actual)
		})
	}
}
