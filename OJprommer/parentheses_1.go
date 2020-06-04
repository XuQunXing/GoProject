package main

import (
	"fmt"
	"strings"
)

/*
45. Parentheses
 */

func main() {
	for  {
		//输入
		var bracket string
		_,err:=fmt.Scanln(&bracket)
		if err!=nil {
			break
		}
		//处理       另一种思路：直接用空字符串替换匹配的括号
		for strings.Contains(bracket,"()"){
			bracket=strings.Replace(bracket,"()","",-1)
		}
		fmt.Println(strings.Count(bracket,"(")+strings.Count(bracket,")"))
	}
}
