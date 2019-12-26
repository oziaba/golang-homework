package main

import "fmt"

func multiply(x int, y int) int {
	res := 0
	for i := 0; i < y; i++{    
		res = res + x
    }
	return res
}

func main() {
	fmt.Println(multiply(20, 3))
}
