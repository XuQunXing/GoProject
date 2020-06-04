package main

import (
	"fmt"
	"sort"
)

/*
28. Candy
*/

func main() {
	for {
		//输入
		var n int
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		candis := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &candis[i])
		}
		//输出   所有进行异或为0则可以分，去最轻的一个之外求和，不为零则输出no
		ans:=0
		for _,value:=range candis{
			ans^=value
		}
		if ans!=0 {
			fmt.Println("NO")
			continue
		}
		sort.Slice(candis, func(i, j int) bool {
			return candis[i]<candis[j]
		})
		solo:=candis[1:len(candis)]
		total:=0
		for _,value:=range solo{
			total+=value
		}
		fmt.Println(total)
	}
}
