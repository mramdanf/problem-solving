package megesortedarr

import "fmt"

func Merge(nums1 []int, m int, nums2 []int, n int)  {
    j := 0
    for i := m; i < m+n; i++ {
        if len(nums1) == m {
            nums1 = append(nums1, nums2[j])
		} else {
            nums1[i] = nums2[j]
        }
        j++
    }
	nums1 = bubbleSort(nums1)
    fmt.Println(nums1)
}

func bubbleSort(arr []int) []int {
	size := len(arr)
	for step := 0; step < size - 1; step++ {
		for i := 0; i < size - step - 1; i++ {
			if arr[i] > arr[i+1] {
				tmp := arr[i+1]
				arr[i+1] = arr[i]
				arr[i] = tmp
			}
		}
	}
	return arr
}