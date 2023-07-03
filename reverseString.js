function solutionOne(str) {
  const arrStr = str.split('')
	arrStr.reverse()
	return arrStr.join('')
}

function solutionTwo(str) {
	const arrStr = str.split('')
	let result = ''
	for (let char of str) {
		result = char + result
	}
	return result
}

console.log(solutionOne('ramdan')) // nadmar
console.log(solutionOne('Greetings!')) // !sgniteerG
console.log(solutionTwo('ramdan')) // nadmar
console.log(solutionTwo('Greetings!')) // !sgniteerG