// two sums problem

// n^2 approach, brute force. 

package main

/*
Takes in an array nums, and a target int as a parameter. Returns an int array of two indices that add up to the target. 
Cannot be duplicates
*/
func twoSumsBrute(nums []int, target int) []int {
	//arr := []int{}

	for i := 0; i < len(nums); i++ {
		for j:= i+1; j < len(nums); j++ {
			if (nums[i] + nums[j] == target && i != j) {
				arr := []int{i, j}
				return arr
			}
		}
	}
	//return arr
	return nil
}

// O(n)
func twoSumsHash(nums []int, target int) []int {
	m := make(map[int]int)
	for index, num := range nums {
		// index = 0
		// num = 2
		// target = 9
		// other = 9 - 2 = 7 
		// value = index 1 (7)
		// then isFound is true
		// when you access a map with a key, you get two return values
		// the value stored for that key 
		// a boolean tells you whether the key is found in the map. 
		value, isFound := m[target-num]
		//fmt.Println(value)
		// if the remaining number from target-num exists, return a slice 
		if isFound {
			return []int{value, index}
		}
		//
		m[num] = index
	}
	// if no pairs are found return nil
	return nil
}
