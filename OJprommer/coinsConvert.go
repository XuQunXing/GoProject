package main

import "fmt"

/*
56.钱币兑换
*/
func main() {
	//3x+2y+1z=N 计算总共多少种
	for {
		var n, ans int
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			break
		}
		//处理

		three := n / 3
		for i := 0; i <= three; i++ {
			ans+= (n-3*i)/2+1
		}

		fmt.Println(ans)
	}

}
