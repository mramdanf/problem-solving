class Dictionary {
	constructor(words) {
		const wordMap = words.reduce((acc, word) => {
			acc[word] = word
			return acc
		}, {})
		this.dict = wordMap
	}

	isInDict(word) {
		if (this.dict[word]) {
			return true
		}

		const template = word.replaceAll('*', '.')
		const reg = new RegExp(template)
		return Object.values(this.dict).some(w => reg.test(w))
	}
}

const test = new Dictionary(['cat', 'car', 'plan'])

console.log(test.isInDict('cat')) // true
console.log(test.isInDict('car')) // true
console.log(test.isInDict('joe')) // false
console.log(test.isInDict('*at')) // true
console.log(test.isInDict('*a*')) // true
console.log(test.isInDict('*aj*')) // false