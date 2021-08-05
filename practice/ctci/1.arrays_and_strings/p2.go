package main

func isPermutation(word1, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	counter := make([]int, 256)
	for _, c := range word1 {
		counter[c]++
	}
	for _, c := range word2 {
		counter[c]--
	}

	for i := 0; i < 256; i++ {
		if counter[i] != 0 {
			return false
		}
	}

	return true
}

// func main() {
// 	fmt.Println(isPermutation("dog", "god"))
// 	fmt.Println(isPermutation("dog", "God"))
// 	fmt.Println(isPermutation("Dog", "God"))
// 	fmt.Println(isPermutation("doG", "God"))
// }
