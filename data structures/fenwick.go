package main

func update(index, val int, tree *[]int) {
	index++
	for index < len(*tree) {
		(*tree)[index] += val
		index += index & -index
	}
}

func query(index int, tree *[]int) int {
	res := 0
	for index >= 1 {
		res += (*tree)[index]
		index -= index & -index
	}
	return res
}

// func main() {
// 	arr := []int{1, 4, 3, 2, -4, 8}
// 	tree := make([]int, len(arr)+1)

// 	for i, elem := range arr {
// 		update(i, elem, &tree)
// 	}

// 	fmt.Println(query(2, &tree)) //sum of 2 first elems
// 	fmt.Println(query(3, &tree)) //sum of 3 first elems
// 	fmt.Println(query(5, &tree)) //sum of 5 first elems
// }
