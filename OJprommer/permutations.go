package main

import "fmt"

/*
289. 难弄的排列组合
*/
func main() {
	for {
		//输入
		var m, n, k int
		_, err := fmt.Scan(&m, &n, &k)
		if err != nil {
			break
		}
		//处理
		n = n - m*k
		if n < 0 {
			fmt.Println(0)
		} else if n == 0 {
			fmt.Println(1)
		} else {
			s := 1
			n=n+m-1
			for i := 0; i < m-1; i++ {
				s *= n - i
			}
			for i := 2; i <= m-1; i++ {
				s /= i
			}
			fmt.Println(s % 5201314)
		}
	}
}
