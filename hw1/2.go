package main

import "fmt"

func fibonacci1 (n int) int {
	var x = 0
	var y = 1
	var res = 0
	if n == 0 {
		res = 0
	}
	if n == 1 {
		res = 1
	}
	for i := 1; i < n; i++ {
		res = x + y
		x = y
		y = res
	}
	return res
}

func main() {
	fmt.Println(fibonacci1(9))
}
