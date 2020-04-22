package main

import "fmt"

func main() {
	fmt.Println(isValid("(]"))
}

func isValid(s string) bool {
	result := make([]string, 0)
	validmap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	for _, item := range s {
		r := fmt.Sprintf("%c", item)
		cores, _ := validmap[r]
		if len(result) > 0 && result[len(result) - 1] == cores {
			if len(result) == 1 {
				result = make([]string, 0)
			} else {
				result = result[0:len(result) - 1]
			}
		} else {
			result = append(result, r)
		}
	}
	if len(result) == 0 {
		return true
	}
	return false
}