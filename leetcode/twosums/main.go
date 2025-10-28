package main

import (
	"fmt"
)


func main() {
	// test case
	result := twoSumsBrute([]int{2, 7, 11, 15}, 9)
	fmt.Println("Brute: ", result)

	result2 := twoSumsHash([]int{2, 7, 11, 15}, 9)
	fmt.Println("Hashmap: ", result2)
}