package main

import (
	"fmt"
	"time"
)

// O(n^2) time
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
	return arr
}


func main() {
	arr := []int{10,9,8,7,6,5,4,3,2,1}
	fmt.Println(arr)
	start := time.Now()
	fmt.Println(BubbleSort(arr))
	fmt.Println(time.Since(start))
}
