package main

import "fmt"

func unique_count(arr[] int) int {
	
	var uniq []int
	for i := 0; i < len(arr); i++ {
		k := 0
		for j := 0; j < len(uniq); j++ {
			if arr[i] != uniq[j] {
				k++
			}
		}
		if k == len(uniq) {
			uniq = append(uniq, arr[i])
		}
	}
	return len(uniq)
}

func main() {
	var arr = []int{1, 2, 3, 4, 1, 2, 2, 3, 2}
	fmt.Println(unique_count(arr))
}