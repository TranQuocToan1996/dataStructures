package dynamic

// Time On
// Space On for map
func FibonacciTopDown(n int) int64 {
	firstTwoCases := map[int]int64{
		0: 0,
		1: 1,
	}

	return fibonacciCache(n, firstTwoCases)
}

func fibonacciCache(n int, cache map[int]int64) int64 {

	if val, found := cache[n]; found {
		return val
	}

	// if not found, recursive calculate
	cache[n] = fibonacciCache(n-1, cache) + fibonacciCache(n-2, cache)
	return cache[n]

}
