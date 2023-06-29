function countingValleys(steps, path) {
    // Write your code here
    let track = 0;
    let down = false;
    let valleys = 0;
    [...path].forEach(p => {
        if (p === 'U' && down === true) {
            track += 1
        }
        if (p === 'D') {
            track -= 1
            down = true;
        }
        console.log({ track, p, down })
        if (track === 0 && down) {
            down = false;
            valleys += 1;
        }
       
    })
    if (track !== 0) return 0
    return valleys;
}


console.log(countingValleys(8, 'UDDDUDUU'))