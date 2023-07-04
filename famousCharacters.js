function searchFamous(str) {
	const arrStr = str.split('')
	const uniqueChars = new Set(arrStr)
	let result = {}
	uniqueChars.forEach(c => {
		let totalFound = 0
		arrStr.forEach(s => {
			if (c === s) {
				totalFound++
			}
		})
		if (totalFound > 0) {
			result[c] = totalFound
		}
	})
	let maxChar = ''
	let maxFound = 0
	Object.keys(result).forEach(key => {
		if (result[key] > maxFound) {
			maxFound = result[key]
			maxChar = key
		}
	})
	console.log(`${maxChar} total of ${maxFound}`)
}

searchFamous('ramdan') // a total of 2
searchFamous('uuuramdan') // u total is 3
searchFamous('ramdn') // none