package main

import "math"

func sortedSquares(nums []int) []int {
    var result []int
    for i := 0; i < len(nums); i++ {
        result = append(result, int(math.Pow(float64(nums[i]), float64(2))))
    }
    return bubbleSort(result)
}

func bubbleSort(nums []int) []int {
    size := len(nums)
    for step := 0; step < size-1; step++ {
        for i := 0; i < size-step-1; i++ {
            if nums[i] > nums[i+1] {
                tmp := nums[i]
                nums[i] = nums[i+1]
                nums[i+1] = tmp
            }
        }
    }
    return nums
}