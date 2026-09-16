func numberOfSets(n int, k int) int {
	const MOD int64 = 1000000007
	N := n + k - 1
	R := 2 * k
	dp := make([]int64, R+1)
	dp[0] = 1

	for i := 1; i <= N; i++ {
		for j := R; j >= 1; j-- {
			dp[j] = (dp[j] + dp[j-1]) % MOD
		}
	}

	return int(dp[R])
}