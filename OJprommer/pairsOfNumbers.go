package main
/*
\42. Pairs of Numbers
  方法：正向思维 暴力法，复杂度2^n  有答案超出限制
 */
import (
	"fmt"
)
type pair struct {
	a int
	b int
}
func main() {
	for true {
		//输入
		var k int
		_,err:=fmt.Scanf("%d",&k)
		if err!=nil {
			break
		}
		//处理
		count:=0
		freq:=1
		slice:=make([]pair,0,freq)
		slice=append(slice,pair{
			a: 1,
			b: 1,
		})
		//先判断，后加减   优化：只判断上次增加的是否包含   只有两个数
		var flag bool  =false
		for  {
			for _,value:=range slice  {
				if judge(value,k) {
					fmt.Println(count)
					flag=true
					break
				}
			}
			if flag==true {
				break
			}
			freq*=2
			temp:=make([]pair,0,freq)
			for _,value:=range slice{
				temp= append(temp, pair{
					a: value.a+value.b,
					b: value.b,
				})
				temp= append(temp, pair{
					a: value.a,
					b: value.a+value.b,
				})
			}
			slice=temp
			count++
		}

	}
}
//判断是否已经有了
func judge(p pair,target int) bool{
	if p.a==target||p.b==target {
		return true
	}else {
		return false
	}
}
