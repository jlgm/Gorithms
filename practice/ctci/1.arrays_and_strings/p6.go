package main

import "fmt"

func compress(word string) string {
	ans, curCount := "", 0
	last := '&'

	for _, c := range word {
		if c == last {
			curCount++
		} else {
			if last != '&' {
				ans += fmt.Sprintf("%d%c", curCount, last)
			}
			last = c
			curCount = 1
		}
	}

	ans += fmt.Sprintf("%d%c", curCount, last)

	if len(ans) >= len(word) {
		return word
	}
	return ans
}

// func main() {
// 	fmt.Println(compress("abbbbcccccaaa"))
// 	fmt.Println(compress("abcdef"))
// }
