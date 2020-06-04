package main

import "fmt"

/*
55. Life, the Universe, and Everything
 */

func main() {
	for {
		var n int
		_, err := fmt.Scanf("%d", &n)
		if n==42||err != nil {
			break
		}
		//处理
		fmt.Println(n)
	}
}
