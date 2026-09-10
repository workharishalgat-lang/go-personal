package main

import (
	"fmt"
)

// coin change problem

func main() {
	fmt.Println(minCoins([]int{10, 20}, 20))
}

// give the minimum coins required to reach the figure

func minCoins(coins []int, figure int) int {
	dp := make([]int, figure+1)
	dp[0] = 0
	for i := 1; i <= figure; i++ {
		dp[i] = figure + 1
	}
	for i := 1; i <= figure; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], 1+dp[i-coin])

			}

		}
	}
	if dp[figure] == figure+1 {
		return -1
	}
	return dp[figure]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
