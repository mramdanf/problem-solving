package removeelement

import "fmt"

func RemoveElement(nums []int, val int) int {
    updatedLength := len(nums)
	
	for step := 0; step < len(nums); step++ {
		for i := 0; i < updatedLength; i++ {
			// fmt.Print(i, " ")
			if nums[i] == val {
				removeAnywhere(&nums, i, &updatedLength)
			}
		}
	}
	// fmt.Print(updatedLength, " -> ")
	// printArray(nums, updatedLength)
	return updatedLength
}

func removeAnywhere(arr *[]int, index int, length *int) {
	// delete last item
	if index == *length-1 {
		(*arr)[index] = -1
		*length--
		// fmt.Println(*arr, *length)
		return
	}

	// delete first and middle item
	for i := index+1; i < len(*arr); i++ {
		(*arr)[i-1] = (*arr)[i]
	}
	*length = *length - 1
	// fmt.Println(*arr, *length)
}

func printArray(arr []int, lenght int) {
	for i := 0; i < lenght; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}