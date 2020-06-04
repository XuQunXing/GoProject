package main

import "fmt"

/*
82. 猴子爬山
 */

func main() {

	for true {
		//输入
		var n int
		_,err:=fmt.Scanf("%d",&n)
		if err!=nil {
			break
		}
		//处理  3x+y=n
		if n<3 {
			fmt.Println(0)
		}
		three:=n/3
		ans:=1
		for i:=1;i<=three;i++  {
			one:=n-3*i+i
			temp:=one
			down:=i
			dome :=down
			for j:=2;j<=i ;j++  {
				temp-=1
				one*=temp
				dome -=1
				down*= dome
			}
			ans+=one/down
		}
		fmt.Println(ans)
	}
}
