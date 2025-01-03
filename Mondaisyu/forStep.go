package main

import "fmt"

/*
階段の上り方計算
*/
func CalcStepUp(n, a, b, c int) int {

	var dp []int
	dp = append(dp, 1)
	for i := 1; i < n+1; i++ {
		dp = append(dp, 0)

		if i >= c {
			dp[i] = dp[i] + dp[i-c]
		}

		if i >= a {
			dp[i] = dp[i] + dp[i-a]
		}

		if i >= b {
			dp[i] = dp[i] + dp[i-b]
		}
	}

	fmt.Println(dp)
	fmt.Println(a, b, c)
	return dp[n]
}
