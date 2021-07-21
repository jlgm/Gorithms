package main

import "fmt"

// calculates fibonacci in O(logN)
// use a memo (map) for speed
func fib(n int) int {

	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}

	if n&1 == 1 { //odd number
		k := (n + 1) >> 1
		return fib(k)*fib(k) + fib(k-1)*fib(k-1)
	} else {
		k := n >> 1
		return (2*fib(k-1) + fib(k)) * fib(k)
	}
}

func main() {
	fmt.Println(fib(60))
}
