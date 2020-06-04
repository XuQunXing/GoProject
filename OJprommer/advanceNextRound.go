package main

import (
	"fmt"
	"sort"
)

/*
60
*/

func main() {
	for {
		var a, b, count int
		_, err :=fmt.Scanf("%d %d", &a, &b)
		score := make([]int, a, 50)
		for i := 0; i < len(score); i++ {
			 fmt.Scanf("%d", &score[i])
		}
		if err != nil {
			break
		}
		sort.Slice(score, func(i, j int) bool {
			return score[i] > score[j]
		})
		for i := 0; i < len(score); i++ {
			if score[i] != 0 && score[i] >= score[b-1] {
				count++
			}
		}
		fmt.Println(count)
	}
}
