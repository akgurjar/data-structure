package main

import "fmt"

func main() {
	arr := [...]int{1, 12, 3, 6, 11, 2, 9, 5, 13, 8}
	for i := len(arr) - 1; i > 0; i-- {
		swap := false
		fmt.Printf("\n Swapping Round : %v", i)
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
	fmt.Printf("\nSorted Array: %v", arr)
}
