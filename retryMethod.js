async function mockFetch(url) {
  return new Promise((resolve, reject) => {
    const randomDelay = Math.random() * 1000
    setTimeout(() => {
      reject('fail')
    }, randomDelay)
  })
}

function retry(fn, times) {
  return async function() {
    try {
      await fn()
    } catch (error) {
      console.log(times, 'fail')
      if (times > 1) {
        retry(fn, times - 1)()
      }
    }
  }
}

const fetchWithDelay = retry(mockFetch, 10)
fetchWithDelay()