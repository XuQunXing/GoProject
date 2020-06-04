package main

import (
	"fmt"
	"sort"
)

/*
54. 整数分解
 */
func main() {
	for {
		var n int
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			break
		}
		//处理
		//找到所有因子
		count:=0
		copy:=n
		factor:=make([]int,n)
		for i:=2;i<=n;i++ {
			for copy%i==0 {
				factor[count]=i
				count++
				copy=copy/i
			}
			if copy==1 {
				break
			}
		}
		slice:=factor[0:count]
		sort.Slice(slice, func(i, j int) bool {
			return slice[i]<slice[j]
		})
		for i:=0;i<len(slice);i++  {
			if i==len(slice)-1 {
				fmt.Printf("%d=%d\n",slice[i],n)

			}else {
				fmt.Print(slice[i],"*")
			}
		}
	}
}
