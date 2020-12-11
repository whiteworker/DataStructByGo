package sort3way

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func nthUglyNumber(n int) int {
    dp := make([]int,n)
    dp[0]=1
    a,b,c :=0,0,0
    for i:=1 ;i<n;i++{
        numA := dp[a]*2
        numB := dp[b]*2
        numC := dp[c]*2
        dp[i]=min( min(numA,numB),numC)
        if dp[i]==numA{
            a++
        }
        if dp[i]==numB{
            b++
        }
        if dp[i]==numC{
            c++
        }
    }
    return dp[n-1]
}
func min(a,b int)int{
    if a>b{
        return b
    } else{
        return a
    }
}