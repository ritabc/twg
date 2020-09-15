package bench

import "sync"

// FibRecursive will calculate the Fibonacci number recursively
func FibRecursive(n int) int {
	if n < 2 {
		return n
	}
	return FibRecursive(n-1) + FibRecursive(n-2)
}

// FibIterative will calculate the Fibonacci number iteratively
func FibIterative(n int) int {
	if n == 0 {
		return 0
	}
	a, b := 0, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

var (
	memo = []int{0, 1}
)

// FibMemo will calculate the Fibonacci number recursively using a pkg level memo variable. This is not threadsafe
func FibMemo(n int) int {
	if len(memo) <= n {
		FibMemo(n - 1)                                 // make sure the memo is filled in up to n-1
		memo = append(memo, FibMemo(n-1)+FibMemo(n-2)) // append the nth value
	}
	return memo[n]
}

var (
	mutex   sync.RWMutex
	memoMap = map[int]int{
		0: 0,
		1: 1,
	}
)

// FibMemoThreadsafe will calculate the Fibonacci number recursively using a pkg level memo variabel and a pkg level mutex. This is threadsafe, but is more confusing than other solutions
func FibMemoThreadsafe(n int) int {
	mutex.RLock()
	if v, ok := memoMap[n]; ok {
		defer mutex.RUnlock()
		return v
	}
	mutex.RUnlock()
	n2 := FibMemoThreadsafe(n - 2)
	n1 := FibMemoThreadsafe(n - 1)
	mutex.Lock()
	memoMap[n] = n1 + n2
	mutex.Unlock()
	return n1 + n2
}
