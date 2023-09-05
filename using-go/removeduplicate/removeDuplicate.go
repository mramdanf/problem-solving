package removeduplicate

import "fmt"

func RemoveDuplicate(nums []int) {
	prev := nums[0]
	l := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			nums[l] = nums[i]
			l++
		}
		prev = nums[i]
	}
	printArray(nums, l)
}

func printArray(nums []int, length int) {
	for i := 0; i < length; i++ {
		fmt.Print(nums[i], " ")
	}
	fmt.Println()
}