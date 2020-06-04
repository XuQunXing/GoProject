package main
/*
\42. Pairs of Numbers
  方法：正向思维 暴力法，复杂度2^n  有答案超出限制
 */
import (
	"fmt"
)


func main() {
	var n,rec int
	for true {
		//输入
		_,err:=fmt.Scanf("%d",&n)
		if err!=nil {
			break
		}
		//处理
		if n==1 {
			fmt.Println(0)
			continue
		}
		var ans int=n
		for i:=1;i<=n ;i++  {
			rec=0
			rec=dfs(n,i,n,rec)
			if ans>rec {
				ans=rec
			}
		}
		fmt.Println(ans)
	}
}
//dfs  深度优先
func dfs(x int,y int,n int,rec int) int{
	if y==0 {
		return n
	}
	if y==1 {
		rec+=x-1
		return rec
	}
	rec+=x/y
	return dfs(y,x%y,n,rec)
}
