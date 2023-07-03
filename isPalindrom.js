function isPalindrom(str) {
  const arrStr = str.split('')
	arrStr.reverse()
	const reversedStr = arrStr.join('')
	return str === reversedStr
}

console.log(isPalindrom('ramdan')) // false
console.log(isPalindrom('ovo')) // true
console.log(isPalindrom('rotator')) // true