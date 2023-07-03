function binarySearchIterate(target, arr) {
  let leftIndex = 0;
	let rightIndex = arr.length - 1

	while(leftIndex <= rightIndex) {
		const middle = Math.floor((leftIndex + rightIndex) / 2)
		const val = arr[middle]
		if (target === val) {
			return `Found target at index ${middle}`
		} else if (target > val) {
			leftIndex += 1
		} else if (target < val) {
			rightIndex -= 1
		}
	}

	return `Target not found`
}

const sortedArr = [1, 2, 3, 4, 5, 6]
console.log(binarySearchIterate(5, sortedArr)) // Found target at index 4
console.log(binarySearchIterate(1, sortedArr)) // Found target at index 0
console.log(binarySearchIterate(4, sortedArr)) // Found target at index 3
console.log(binarySearchIterate(10, sortedArr)) // Target not found