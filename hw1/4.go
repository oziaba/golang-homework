package main

import "fmt"

func bubble_sort(arr[] float64){
	var k float64
	for i:=1 ; i < len(arr) ; i++ {
		for j:=0 ; j < len(arr)-i ; j++ {
			if arr[j] > arr[j+1] {
				k = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = k
			}
		}
	}
}

func main() {
	var arr = []float64{1,20,0,4,8}
	bubble_sort(arr)
	fmt.Println(arr)
}