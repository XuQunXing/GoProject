package main

import (
	"fmt"
	"strconv"
)

/*
32. 整数拆分 Ver.2
*/

func main() {
	for {
		//输入
		var n int
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		//输出
		ans:=strconv.Itoa(n)+"="
		inteSpli(n,ans,1)

	}
}
//拆分打印
func inteSpli(n int,s string,limit int)  {
	//拆分玩打印
	if n==0 {
		fmt.Println(s[:len(s)-1])
	}
	//从limit开始，后面值不小于前面
	for i:=limit;i<=n ;i++  {
		s1:=s+strconv.Itoa(i)+"+"
		inteSpli(n-i,s1,i)
	}
}
