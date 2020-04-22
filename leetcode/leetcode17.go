package main

import "fmt"

func main() {
	fmt.Println(letterCombinations(""))
}

func letterCombinations(digits string) []string {
	finals := make([]string, 0)
	combine(&finals, "", digits)
	return finals
}

func combine(finals *[]string, prefix string, letters string) {
	digitLetterMap := map[string][]string {
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	if len(letters) == 0 {
		return
	} else if len(letters) == 1 {
		digits, ok := digitLetterMap[letters]
		if !ok {
			return
		}
		for _, digit := range digits {
			*finals = append(*finals, prefix + digit)
		}
	} else {
		first := letters[0:1]

		digits, ok := digitLetterMap[first]
		if !ok {
			return
		}
		for _, digit := range digits {
			combine(finals, prefix + digit, letters[1:])
		}
	}
}