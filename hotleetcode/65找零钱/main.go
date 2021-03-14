package main
//动态规划（不压缩空间）
func coinChange(coins []int, amount int) int {
    dp:=make([][]int,len(coins)+1)//定义状态
    for i:=0;i<=len(coins);i++{
        dp[i]=make([]int,amount+1)
    }
    //初始条件
    for j:=0;j<=amount;j++{
        dp[0][j]=amount+1
    }
    dp[0][0]=0
    for i:=1;i<=len(coins);i++{
        for j:=0;j<=amount;j++{
            //状态转移方程
            if j >= coins[i-1]{
                dp[i][j]=int(math.Min(float64(dp[i-1][j]),float64(dp[i][j-coins[i-1]]+1)))
            }else{
                dp[i][j]=dp[i-1][j]
            }
        }
    }
    if dp[len(coins)][amount]>amount{
        return -1
    }else{
        return dp[len(coins)][amount]
    }
}
func change(amount int, coins []int) int {
	var dp = make([]int, amount+1)
	dp[0] = 1
	for _, v := range coins {
		for i := 1; i < amount+1; i++ {
			if i-v >= 0 {
				dp[i] += dp[i-v]
			}
		}
	}
	return dp[amount]
}

