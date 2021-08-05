package main

func isUnique(word string) bool {
	m := make([]bool, 26)
	for _, c := range word {
		if m[c-'a'] {
			return false
		}
		m[c-'a'] = true
	}
	return true
}

func isUniqueBitmask(word string) bool {
	m := 0
	for _, c := range word {
		pos := c - 'a'
		if (m>>pos)&1 == 1 {
			return false
		}
		m |= (1 << pos)
	}
	return true
}

// func main() {
// 	fmt.Println(isUniqueBitmask("carlos"))
// 	fmt.Println(isUniqueBitmask("joaizn"))
// 	fmt.Println(isUniqueBitmask("ana"))
// 	fmt.Println(isUniqueBitmask("c"))
// 	fmt.Println(isUniqueBitmask("cc"))
// 	fmt.Println(isUniqueBitmask("joao"))
// }
