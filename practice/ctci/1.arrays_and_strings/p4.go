package main

func isPalinPerm(word string) bool {

	counter := make([]int, 26)
	for _, c := range word {
		counter[c-'a']++
	}

	odd := 0
	for i := 0; i < 26; i++ {
		if counter[i]&1 == 1 {
			odd++
		}
	}

	if len(word)&1 == 0 {
		return odd == 0
	} else {
		return odd == 1
	}
}

// func main() {
// 	fmt.Println(isPalinPerm("bacbaf"))
// 	fmt.Println(isPalinPerm("aaaan"))
// }
