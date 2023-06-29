class Dictionary {
	constructor(words) {
		this.dict = words
	}

	isInDict(word) {
		const template = word.replaceAll('*', '.')
		const reg = new RegExp(template)
		return this.dict.some(w => reg.test(w))
	}
}

const test = new Dictionary(['cat', 'car', 'plan'])

console.log(test.isInDict('cat')) // true
console.log(test.isInDict('car')) // true
console.log(test.isInDict('joe')) // false
console.log(test.isInDict('*at')) // true
console.log(test.isInDict('*a*')) // true
console.log(test.isInDict('*aj*')) // false