/*
Problem: https://leetcode.com/problems/valid-parentheses/
Author: Thada Wangthammang
*/

package ValidParentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name     string
	input    string
	expected bool
}{
	{"A Pair of {}", "{}", true},
	{"Single {", "{", false},

	{"Single (", "(", false},
	{"A Pair of ()", "()", true},

	{"Not pair {)", "{)", false},
	{"Not pair (}", "(}", false},
}

func TestTwoSum(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(subtest *testing.T) {
			// subtest.Parallel()
			actual := isValid(test.input)
			assert.Equal(subtest, actual, test.expected)
		})
	}
}
