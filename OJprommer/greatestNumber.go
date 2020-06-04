package main

import (
	"fmt"
	"sort"
)

/*
58. Greatest Number
*/

func main() {
	i:=1
	for {
		//输入
		var n, m int
		fmt.Scanf("%d %d\n", &n, &m)
		if n == 0 && m == 0 {
			break
		}
		dict := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scanln(&dict[i])
		}
		fmt.Scanln()
		//处理
		//先求任意两个数的和
		temp := make([]int, n*n)
		index := 0
		for _, value := range dict {
			for _, value1 := range dict {
				temp[index] = value + value1
				index++
			}
		}
		//排序
		sort.Slice(temp, func(i, j int) bool {
			return temp[i] < temp[j]
		})
		//二分查找
		left := 0
		right := len(temp) - 1
		max := 0
		for left <= right {
			addNum := temp[left] + temp[right]
			if addNum > m {
				right--
			} else {
				if addNum > max {
					max = addNum
				}
				left++
			}
		}
		fmt.Printf("Case %d: %d\n", i, max)
		fmt.Println()
		i++
	}
}
