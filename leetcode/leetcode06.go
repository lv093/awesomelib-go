package main

import (
)

func main() {
	println(convert("ab", 1))
}

func convert(s string, numRows int) string {

	headStep := 2*numRows - 2
	line := 0
	sLen := len(s)
	res := ""
	if sLen <= numRows || numRows == 1 {
		return s
	}
	for line < numRows {
		i := line
		if line != 0 && line != numRows -1 {
			res = res + s[line: line+1]
		}
		oddNum := headStep - line
		unevenNum := headStep + i
		for i < sLen {

			if line == 0 || line == numRows - 1 {
				res = res + s[i:i+1]
				i = headStep + i
			} else {
				i = oddNum
				if i < sLen {
					res = res + s[i:i+1]
				}
				i = unevenNum
				if i < sLen {
					res = res + s[i:i+1]
				}
				oddNum = oddNum + headStep
				unevenNum = unevenNum + headStep
			}
		}
		line++
	}
	return res
}