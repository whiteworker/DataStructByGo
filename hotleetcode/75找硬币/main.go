func coinChange(coins []int, amount int) int {
    if len(coins)==0 || amount==0{
        return 0
    }
    dp:=make([]int,amount+1)
    for i:=1;i<=amount;i++{
        dp[i]=	math.MaxInt64
        for e:=range coins{
            if i-coins[e]>=0 && dp[i-coins[e]]!=math.MaxInt64{
                dp[i]=min(dp[i-coins[e]]+1,dp[i])
            }
        }
    }
    if dp[amount]==math.MaxInt64{
        return -1
    }
    return dp[amount]
}

func min(a,b int) int{
    if a>b{
        return b
    }
    return a
}
