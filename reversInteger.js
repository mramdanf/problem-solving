function reverseInt(integerNumber) {
  const reversedNum = parseInt(integerNumber.toString().split('').reverse().join(''))
  if (integerNumber < 0) {
    return  reversedNum * -1
  }

  return reversedNum
}

console.log(reverseInt(2031)) // 1302
console.log(reverseInt(1213)) // 3121
console.log(reverseInt(-54)) // -45