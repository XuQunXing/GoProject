package main
/*
31. 整数拆分
 */
import "fmt"

/*
31
 */
func main() {
	var a int
	for true {
		_,err:=fmt.Scanf("%d",&a)
		if err!=nil {
			break
		}
		fmt.Println(recursion(a))
	}
}
//递归
func recursion(n int) int  {
	//二维切片
	dp:=make([][]int,n+1)
	for i:=range dp{
		dp[i]=make([]int,n+1)
	}

	//dp
	for i:=1;i<n+1 ;i++  {
		for j:=1;j<n+1 ;j++  {
			if i==1||j==1 {
				dp[i][j]=1
			}else if i==j {
				dp[i][j]=1+dp[i][j-1]
			}else if i<j {
				dp[i][j]=dp[i][i]
			}else {
				dp[i][j]=dp[i-j][j]+dp[i][j-1]
			}
		}
	}
	return dp[n][n]
}