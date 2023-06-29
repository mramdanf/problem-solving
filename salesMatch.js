function sockMerchant(n, ar) {
    // Write your code here
    let match = 0
    ar.forEach((a, i) => {
        let couple = 0;
        [...ar.slice(i+1)].forEach(b => {
            // console.log({ a, b })
            if (a === b) {
                couple += 1
            }
        })
        console.log({ a, couple })
        if (couple > 0) {
            match += (couple % 2)
        }
        
    })
    return match
}

console.log(sockMerchant(2, [10,20,20,10,10,30,50,10,20]))