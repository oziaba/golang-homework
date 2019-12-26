package main

import "fmt"

func fibonacci2 (n int) int {
	var res = 0
	if n == 0 {
		res = 0
	}
	if n == 1 {
		res = 1
	}
	for i := 1; i < n; i++ {
		res = fibonacci2(n - 1) + fibonacci2(n - 2)
	}
	return res
}

func main() {
	fmt.Println(fibonacci2(5))
}
