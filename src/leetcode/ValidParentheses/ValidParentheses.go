/*
Problem: https://leetcode.com/problems/valid-parentheses/
Author: Thada Wangthammang
*/
package ValidParentheses

// import "github.com/emirpasic/gods/stacks/arraystack"

func isValid(s string) bool {
	return len(s)%2 == 0
}
