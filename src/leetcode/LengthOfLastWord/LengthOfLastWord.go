/*
Problem: https://leetcode.com/problems/
Author: Thada Wangthammang
Speed:
*/

package LengthOfLastWord

import "strings"

func lengthOfLastWord(s string) int {
	trimString := strings.TrimSpace(s)
	return len(getLastWord(trimString))
}

func getLastWord(s string) string {
	words := strings.Split(s, " ")
	return words[len(words)-1:][0]
}
