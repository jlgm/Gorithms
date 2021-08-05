package main

func URLify(word []rune) string {
	for i, c := range word {
		if c == ' ' {
			copy(word[i+3:], word[i+1:])
			word[i] = '%'
			word[i+1] = '2'
			word[i+2] = '0'
		}
	}
	return string(word)
}

// func main() {
// 	str := "mr john smith----"
// 	fmt.Println(URLify([]rune(str)))
// }
