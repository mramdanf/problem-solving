function searchFamous(str) {
	let result = {}
	for(let char of str) {
		result[char] ? result[char]++ : result[char] = 1
	}
	let maxNum = 0
	let maxChar = 0
	for (let char in result) {
		if (result[char] > maxNum) {
			maxNum = result[char]
			maxChar = char
		}
	}
	console.log(`${maxChar} total of ${maxNum}`)
}

searchFamous('ramdan') // a total of 2
searchFamous('uuuramdan') // u total is 3
searchFamous('ramdn') // none