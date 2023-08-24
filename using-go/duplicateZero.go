package main

import "fmt"

func duplicateZeros(arr []int)  {
	for i := 0; i < len(arr); i++ {
		inserted := false
		if arr[i] == 0 && !inserted {
			insertZeroAt(&arr, i+1)
			inserted = true
		} else {
			inserted = false
		}
	}
	// fmt.Println(arr)
}

func insertZeroAt(arr *[]int, index int) {
    *arr = append(*arr, -1)
	fmt.Println(*arr, index, len(*arr))
    for i := len(*arr) - 1; i > index; i-- {
		// fmt.Println(i, (*arr)[i], (*arr)[i-1])
        (*arr)[i] = (*arr)[i-1]
    }
	fmt.Println(*arr)
    (*arr)[index] = 0
}