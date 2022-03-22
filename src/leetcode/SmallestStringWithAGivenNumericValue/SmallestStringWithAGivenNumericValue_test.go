package SmallestStringWithAGivenNumericValue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name     string
	n        int
	k        int
	expected string
}{
	{"Example in Problem", 3, 27, "aay"},
}

func TestSmallestStringWithAGivenNumericValue(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(subtest *testing.T) {
			// subtest.Parallel()
			actual := getSmallestString(test.n, test.k)
			assert.Equal(subtest, test.expected, actual)
		})
	}
}
