/*
lower_bound is the index of first element which is not LESS than val (1, 2, [3], 3, 3, 4, 5) (val=3)
upper_bound is the index of first element which is GREATER than val (1, 2, 3, 3, 3, [4], 5) (val=3)
*/
package main

import "fmt"

func lowerBound(arr *[]int, elem int) int {
	lo, hi := 0, len(*arr)-1
	for lo <= hi {
		mid := (lo + hi) >> 1
		if (*arr)[mid] < elem {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func upperBound(arr *[]int, elem int) int {
	lo, hi := 0, len(*arr)-1
	for lo <= hi {
		mid := (lo + hi) >> 1
		if (*arr)[mid] <= elem {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func main() {
	arr := []int{1, 2, 3, 3, 3, 5, 5, 5, 5, 5, 5} // array has to be sorted

	fmt.Println("3:", lowerBound(&arr, 3), upperBound(&arr, 3))

	fmt.Println("4:", lowerBound(&arr, 4), upperBound(&arr, 4))

	fmt.Println("5:", lowerBound(&arr, 5), upperBound(&arr, 5))
}
