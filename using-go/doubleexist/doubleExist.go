package doubleexist

func CheckDoubleExist(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		// var found bool
		for j := 0; j < len(arr); j++ {
			
			if i == j {
				continue
			}

			if arr[i] == (2*arr[j]) {
				// fmt.Println(arr[i], arr[j])
				return true
			}
		}
		// if found {
		// 	return true
		// }
	}
	return false
}