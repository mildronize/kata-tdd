/*
Problem: https://leetcode.com/problems/valid-parentheses/
Author: Thada Wangthammang
Speed: O(n)
Difficulty: Easy
*/

// Parenthesis = Bracket

package ValidParentheses

import "github.com/emirpasic/gods/stacks/arraystack"

func isValid(s string) bool {
	stack := arraystack.New()
	for _, ch := range s {
		if isOpenParenthesis(ch) {
			stack.Push(ch)
		} else if isCloseParenthesis(ch) {
			recentValue, _ := stack.Peek()
			if recentValue == '{' && ch == '}' {
				stack.Pop()
			} else if recentValue == '(' && ch == ')' {
				stack.Pop()
			} else if recentValue == '[' && ch == ']' {
				stack.Pop()
			} else {
				return false
			}
		}
	}
	return stack.Size() == 0
}

func isOpenParenthesis(ch rune) bool {
	return ch == '{' || ch == '(' || ch == '['
}

func isCloseParenthesis(ch rune) bool {
	return ch == '}' || ch == ')' || ch == ']'
}
