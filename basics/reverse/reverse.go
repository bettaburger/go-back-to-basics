package reverse 

import (
	"fmt"
)

// reverse an integer array
func R1(nums []int) []int {
	fmt.Println("Reversing an integer array...")

	// reversed list. 
	output := make([]int, len(nums))
	//starting index for reversed list
	start := 0

	for i := len(nums) - 1; i >= 0; i-- {
		// fill in array
		output[start] = nums[i]
		// increment start to the next index. 
		start++ 
	}
	return output 
}
// reverse a string array

// reverse a linkedlist
