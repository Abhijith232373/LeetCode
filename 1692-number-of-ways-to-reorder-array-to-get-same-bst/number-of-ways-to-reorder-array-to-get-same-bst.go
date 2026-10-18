func numOfWays(nums []int) int {
    const MOD = 1000000007
    n := len(nums)
    comb := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        comb[i] = make([]int, n+1)
        comb[i][0] = 1
        comb[i][i] = 1
        for j := 1; j < i; j++ {
            comb[i][j] = (comb[i-1][j-1] + comb[i-1][j]) % MOD
        }
    }
    var solve func([]int) int
    solve = func(arr []int) int {
        if len(arr) <= 2 {
            return 1
        }
        root := arr[0]
        left := []int{}
        right := []int{}
        for _, v := range arr[1:] {
            if v < root {
                left = append(left, v)
            } else {
                right = append(right, v)
            }
        }
        leftWays := solve(left)
        rightWays := solve(right)
        ways := comb[len(left)+len(right)][len(left)]
        ways = (ways * leftWays) % MOD
        ways = (ways * rightWays) % MOD
        return ways
    }
    return (solve(nums) - 1 + MOD) % MOD
}