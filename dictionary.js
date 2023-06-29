class Dictionary {
	constructor(words) {
		this.dict = new Set(words)
	}

	isInDict(word) {
		return this.dict.has(word)
	}
}

const test = new Dictionary(['cat', 'car', 'plan'])

console.log(test.isInDict('cat')) // true
console.log(test.isInDict('car')) // true
console.log(test.isInDict('joe')) // false