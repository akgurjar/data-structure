package main

import "fmt"

func main() {
	num := 12345
	temp := num
	rev := 0
	for temp > 0 {
		rev = rev*10 + temp%10
		temp = temp / 10
	}
	fmt.Printf("Reverse : %v \n", rev)
}
