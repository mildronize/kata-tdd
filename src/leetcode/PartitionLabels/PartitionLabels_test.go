package PartitionLabels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name     string
	input    string
	expected []int
}{
	{"simple", "a", []int{1}},
	{"simple", "ababcbacad", []int{9, 1}},
	{"simple", "ababcbacadedf", []int{9, 3, 1}},
	{"simple", "ababcbacadefegde", []int{9, 7}},
	{"simple", "ababcbacadefegdehijhklij", []int{9, 7, 8}},
	{"simple", "eccbbbbdec", []int{10}},
}

func TestPartitionLabels(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(subtest *testing.T) {
			// subtest.Parallel()
			actual := partitionLabels(test.input)
			assert.ElementsMatch(subtest, test.expected, actual)
		})
	}
}

var findFirstLetterIndexTests = []struct {
	name     string
	str      string
	findChar rune
	expected int
}{
	{"Return 0", "aaaa", 'a', 0},
	{"Return 0", "bbabba", 'a', 2},
	{"Return -1 when not found", "aaaa", 'b', -1},
}

func TestFindFirstLetterIndex(t *testing.T) {
	for _, test := range findFirstLetterIndexTests {
		t.Run(test.name, func(subtest *testing.T) {
			actual := findFirstLetterIndex(test.str, test.findChar)
			assert.Equal(subtest, test.expected, actual)
		})
	}
}

var findLastLetterIndexTests = []struct {
	name     string
	str      string
	findChar rune
	expected int
}{
	{"Return 0", "aaaa", 'a', 3},
	{"Return 0", "bbabba", 'a', 5},
	{"Return -1 when not found", "aaaa", 'b', -1},
}

func TestLastLetterIndex(t *testing.T) {
	for _, test := range findLastLetterIndexTests {
		t.Run(test.name, func(subtest *testing.T) {
			actual := findLastLetterIndex(test.str, test.findChar)
			assert.Equal(subtest, test.expected, actual)
		})
	}
}
