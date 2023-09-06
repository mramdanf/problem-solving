package replaceelementwithhighest

func ReplaceElement(arr []int) []int {
	if len(arr) == 1 {
		arr[0] = -1
		return arr
	}

	for i := 0; i < len(arr); i++ {
		if i == len(arr) - 1 {
			arr[i] = -1
			continue
		}

		heighest := arr[i+1]
		for j := i+2; j < len(arr); j++ {
			if arr[j] > heighest {
				heighest = arr[j]
			}
		}
		arr[i] = heighest
	}

	return arr
}