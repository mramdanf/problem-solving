class Dictionary {
	constructor(words) {
		const wordMap = words.reduce((acc, word) => {
			acc[word] = true
			word.split('').forEach((_, i) => {
				const start = word.slice(0, i)
				const end = word.slice(i + 1)
				const partialWord = `${start}*${end}`
				acc[partialWord] = true
			})
			return acc
		}, {})
		this.dict = wordMap
	}

	isInDict(word) {
		return !!this.dict[word]
	}
}

const test = new Dictionary(['cat', 'car', 'plan'])

console.log(test.isInDict('cat')) // true
console.log(test.isInDict('*at')) // true
console.log(test.isInDict('*a*')) // false
console.log(test.isInDict('don')) // false