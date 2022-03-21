/*
Problem: https://leetcode.com/problems/partition-labels/
Author: Thada Wangthammang
Speed: O(n)
Difficulty: Medium
Note: Daily Problem
Tags: substring, greedy
*/

package PartitionLabels

type Letter struct {
	first int
	last  int
}

var aLetterLowercaseAscii int = 97

func partitionLabels(s string) []int {
	var blocks []int = []int{}
	var currentLastIndex int = 0
	var blockSize int = 0
	var letters [26]Letter
	// Prepare metadata for processing string
	for i := 0; i < len(letters); i++ {
		letters[i].first = findFirstLetterIndex(s, rune(aLetterLowercaseAscii+i))
		letters[i].last = findLastLetterIndex(s, rune(aLetterLowercaseAscii+i))
		// fmt.Printf("%c: %d, %d\n", rune(aLetterLowercaseAscii+i), letters[i].first, letters[i].last)
	}
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		// Starting new block
		if blockSize == 0 {
			// find letter in letters DB
			letterIndex := runeToIndex(runes[i])
			// init currentLastIndex
			currentLastIndex = letters[letterIndex].last
		}

		// Last Loop, add new block finally.
		if i+1 == len(runes) {
			blocks = append(blocks, blockSize+1)
		} else {
			// Checking next char is should be ending of the block or not.
			nextLetterIndex := runeToIndex(runes[i+1])
			nextFirstIndex := letters[nextLetterIndex].first
			nextLastIndex := letters[nextLetterIndex].last

			// Check if the next char is already found in current block, set a new currentLastIndex
			if nextFirstIndex < currentLastIndex && nextLastIndex > currentLastIndex {
				currentLastIndex = nextLastIndex
			}
			if nextLastIndex <= currentLastIndex {
				blockSize++
			} else {
				blocks = append(blocks, blockSize+1)
				blockSize = 0
			}
		}

	}
	return blocks
}

func runeToIndex(r rune) int {
	return int(r) - aLetterLowercaseAscii
}

func findFirstLetterIndex(s string, findChar rune) int {
	for i, value := range s {
		if findChar == value {
			return i
		}
	}
	return -1
}

func findLastLetterIndex(s string, findChar rune) int {
	runes := []rune(s)
	for i := len(s) - 1; i >= 0; i-- {
		if findChar == runes[i] {
			return i
		}
	}
	return -1
}
