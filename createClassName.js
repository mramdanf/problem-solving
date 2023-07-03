function createClassName(...args) {
	const classes = []
	for (let i = 0; i < args.length; i++) {
		let arg = args[i]
		if (typeof arg === 'string' || (typeof arg === 'number' && arg > 0)) {
			classes.push(arg)
		} else if (typeof arg === 'object' && !Array.isArray(arg) && !!arg) {
			Object.keys(arg).forEach(key => {
					if (arg[key]) {
							classes.push(key)
					}
			})
		} else if (Array.isArray(arg)) {
			return createClassName.apply(null, arg)
		}
	}
	return classes.join(' ')
}

console.log(createClassName('classA', 'classB')) // classA classB
console.log(createClassName('classA', 1)) // classA 1
console.log(createClassName('classA', 0, 1, null, undefined)) // classA 1
console.log(createClassName({ classA: true, classB: true, classC: true })) // classA classB classC
console.log(createClassName([{ classA: true, classB: true, classC: true }, 'classD', 1, null, undefined, 0])) // classA classB classC classD 1
