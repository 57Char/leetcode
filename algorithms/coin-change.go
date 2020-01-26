package algorithms

// https://leetcode.com/problems/coin-change/

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, c := range coins {
			if i-c >= 0 {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func coinChangeV2(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	max := amount + 1
	dp := make([]int, max)
	for i := range dp {
		dp[i] = max
	}
	dp[0] = 0
	for i:=1; i<=amount; i+=1 {
		for j:=0; j<len(coins); j+=1 {
			if coins[j] <= i {
				if dp[i] > dp[i-coins[j]] + 1 {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
	}
	if dp[amount] == max {
		return -1
	}
	return dp[amount]
}