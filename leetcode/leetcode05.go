package main

import "time"

func main() {
	println("hello"[0:1])
	t := time.Now()
	println(int(t.Weekday()))

	println(longestPalindrome("dbac"))
	println(forceGetLongestPalindrome("cbbd"))
}

func longestPalindrome(s string) string {
	n := len(s)
	strPosMap := make(map[int32]int, n)
	for i, item := range s {
		if prev, ok := strPosMap[item]; ok {
			return s[prev:i+1]
		}
		strPosMap[item] = i
	}

	return s[0:1]
}

//暴力破解法
//暴力求解，列举所有的子串，判断是否为回文串，保存最长的回文串。
func forceGetLongestPalindrome(s string) string {
	m := len(s)
	if m <= 1 {
		return s
	}
	max := 1
	longest := s[0:1]
	for i := 0; i < m; i++ {
		for j := i+1; j < m; j++ {
			ok := checkPalindrome(s[i: j+1])
			println(s[i:j+1], ok)
			if ok && j+1-i > max {
				max = j - i
				longest = s[i:j+1]
			}
		}
	}
	return longest
}

func checkPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	right, left := 0, len(s) - 1
	for right < left {
		if s[right] != s[left] {
			return false
		}
		right++
		left--
	}
	return true
}

