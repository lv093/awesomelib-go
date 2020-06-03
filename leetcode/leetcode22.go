package main

import (
	"fmt"
)

/**
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
示例：

输入：n = 3
输出：[
       "((()))",
       "(()())",
       "(())()",
       "()(())",
       "()()()"
     ]
 */
func main() {
	fmt.Println(generateParenthesis(1))
}

var res []string

func generateParenthesis(n int) []string {
	res = make([]string, 0)
	getNextBracket( "(", n-1, n, 1)

	return res
}

/**
方法1：使用栈思想的方式
leftStack 表示左括号集合
rightStack表示右括号集合
 */
func getNextBracket(curr string, left int, right int, free int) (string) {
	if left == 0 && right == 0 {
		res = append(res, curr)
		return curr
	}
	if free == 0 {
		curr = curr + "("
		curr = getNextBracket( curr, left-1, right, free+1)
	} else {
		curr1 := curr + ")"
		curr1 = getNextBracket(curr1, left, right - 1, free-1)
		if left > 0 {
			curr2 := curr + "("
			curr2 = getNextBracket( curr2, left - 1, right, free+1)
		}
	}
	return ""
}
