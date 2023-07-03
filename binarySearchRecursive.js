function binarySearch(target, arr, start, end) {

	if (start > end) {
		return `Target not found`
	}
  
	const middle = Math.floor((start + end) / 2)
	
	if (arr[middle] === target) {
		return `Found target at index ${middle}`
	}

	if (target > arr[middle]) {
		return binarySearch(target, arr, middle + 1, end)
	}

	if (target < arr[middle]) {
		return binarySearch(target, arr, start, middle - 1)
	}
}

function findElement(target, sortedArr) {
	return binarySearch(target, sortedArr, 0, sortedArr.length - 1)
}

const sortedArr = [1, 2, 3, 4, 5, 6]
console.log(findElement(5, sortedArr)) // Found target at index 4
console.log(findElement(1, sortedArr)) // Found target at index 0
console.log(findElement(4, sortedArr)) // Found target at index 3
console.log(findElement(10, sortedArr)) // Target not found