function sum(...nums) {
    console.log(nums.reduce((acc, num) => acc + num, 0))
}

sum(1, 2, 3)