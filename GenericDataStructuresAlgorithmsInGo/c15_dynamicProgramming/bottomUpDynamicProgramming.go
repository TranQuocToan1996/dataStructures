package dynamic

// Time On
// Space O1

func FibonacciBottomUp(n int) int64 {
	table := []int64{0, 1}
	for i := 2; i <= n; i++ {
		table = append(table, table[i-1]+
			table[i-2])
	}
	return table[n]
}

func FibRecursive(n int) int64 {
	if n < 2 {
		return int64(n)
	}
	return FibRecursive(n-1) + FibRecursive(n-2)
}
