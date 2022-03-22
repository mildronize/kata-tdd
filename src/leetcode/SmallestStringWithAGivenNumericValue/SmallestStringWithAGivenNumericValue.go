/*
Problem: https://leetcode.com/problems/smallest-string-with-a-given-numeric-value/
Author: Thada Wangthammang
Speed: O(n) (Can be O(1))
Difficulty: Medium
Tags: Greedy,
Note: Daily Problem
*/

package SmallestStringWithAGivenNumericValue

// Cannot import this lib in leet code
// import "github.com/huandu/xstrings"

var maxAlphabetOrder int = 26
var aLetterLowercaseAscii int = 97

func getSmallestString(n int, k int) string {
	// alphabetical Order start from 1 - 26 (a-z)
	list := make([]rune, n)
	for index := 0; index < n; index++ {
		remain_letters := n - index - 1
		// Ex: n=5, k=73
		// First Iteration
		// list: _____ -> zaaaa (Smallest possible) = 26 + 1 + 1 + 1 < 73
		// then place 'z' and move to next letter
		// so: list will be z____
		if maxAlphabetOrder+remain_letters < k {
			list[remain_letters] = alphabetOrderToRune(maxAlphabetOrder)
			k -= maxAlphabetOrder
		} else {
			// find the Largest alphabet order
			largest := k - remain_letters
			list[remain_letters] = alphabetOrderToRune(largest)
			k -= largest
		}
	}
	return string(list)
}

func alphabetOrderToRune(order int) rune {
	return rune(aLetterLowercaseAscii + order - 1)
}

// func reverseString(s string) string {
// 	runes := []rune(s)
// 	var result []rune
// 	for i := len(runes) - 1; i >= 0; i-- {
// 		result = append(result, runes[i])
// 	}
// 	return string(result)
// }
