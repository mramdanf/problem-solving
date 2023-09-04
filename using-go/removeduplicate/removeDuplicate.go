package removeduplicate

import "fmt"

func RemoveDuplicate(nums []int) {
	updatedLen := len(nums)
	for i := 0; i < updatedLen; {
		if nums[i] == nums[i+1] {
			// delete next duplicate
			for j := i+2; j < updatedLen; j++ {
				nums[j-1] = nums[j]
			}
			
			// if there is deleting at least one
			// than we should decrement the length
			if i+2 < updatedLen || i+2 == updatedLen {
				updatedLen--
			}
			continue
		}
		i++
	}
	printArray(nums, updatedLen)
}

func printArray(nums []int, length int) {
	for i := 0; i < length; i++ {
		fmt.Print(nums[i], " ")
	}
	fmt.Println()
}