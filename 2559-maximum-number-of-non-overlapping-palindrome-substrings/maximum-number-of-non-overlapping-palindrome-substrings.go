func maxPalindromes(s string, k int) int {
	n := len(s)
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		dp[i+1] = dp[i]
		for length := k; length <= i+1; length++ {
			if length%2 == 0 {
				continue
			}

			start := i - length + 1

			if isPalindrome(s, start, i) {
				dp[i+1] = max(dp[i+1], dp[start]+1)
				break
			}
		}
		for length := k; length <= i+1; length++ {
			if length%2 != 0 {
				continue
			}

			start := i - length + 1

			if isPalindrome(s, start, i) {
				dp[i+1] = max(dp[i+1], dp[start]+1)
				break
			}
		}
	}

	return dp[n]
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}