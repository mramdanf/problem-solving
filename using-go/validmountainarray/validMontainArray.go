package validmountainarray

func ValidMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	peakReached := false
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] < arr[i] && i > 0 {
			peakReached = true
		}

		if arr[i+1] <= arr[i] && !peakReached {
			return false
		}

		if arr[i+1] >= arr[i] && peakReached {
			return false
		}
	}

	return peakReached
}