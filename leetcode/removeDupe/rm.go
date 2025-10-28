package main 
 
import (
	"fmt"
)



func removeDuplicates(nums []int) int {
    // Using two pointers approach
    // i points to the last unique element
    // j iterates through the array
    i := 0
    for j := 1; j < len(nums); j++ {
        if nums[j] != nums[i] {
            i++
            nums[i] = nums[j]
            fmt.Println(i, j)
        }
    }
    fmt.Println(i)

    return i + 1
}


func main (){
	//nums := []int{0,0,1,1,1,2,2,3,3,4}
	//result := removeDuplicates(nums)
	result2 := removeDuplicates([]int{1, 1, 2})
	//fmt.Println(result)
	fmt.Println(result2)
}