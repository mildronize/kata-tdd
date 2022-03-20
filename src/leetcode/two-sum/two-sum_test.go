// Ref: https://medium.com/@pliutau/table-driven-tests-in-go-5d7e230681da

package two_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name     string
	nums     []int
	target   int
	expected []int
}{
	{"2+7=9", []int{2, 7, 11, 15}, 9, []int{0, 1}},
}

// func TestTwoSum(t *testing.T) {
// 	nums := []int{2, 7, 11, 15}
// 	target := 9
// 	expected := []int{0, 1}
// 	assert.ElementsMatch(t, twoSum(nums, target), expected)
// }

func TestTwoSum(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(subtest *testing.T) {
			subtest.Parallel()
			actual := twoSum(test.nums, test.target)
			assert.ElementsMatch(subtest, actual, test.expected)
		})
	}
}
