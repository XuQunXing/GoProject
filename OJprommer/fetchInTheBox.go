package main

import "fmt"

/*
89. 盒中取球
*/

func main() {
	var table =[10000]int{0,0,1,0,1,0,1,0,1,1}
	for i:=10;i<10000 ;i++  {
		if table[i-8]+table[i-7]+table[i-3]+table[i-1]<4 {
			table[i]=1
		}else{
			table[i]=0
		}
	}
	for {
		//输入
		var n int
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			break
		}
		bottles := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &bottles[i])
		}
		//处理 一千以内的数的输赢建表存起来，直接查
		for _,value:=range bottles {
			fmt.Println(table[value])
		}

	}
}
