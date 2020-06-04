package main

import (
	"fmt"
	"strings"
)

/*
69
*/

func main() {

	for true {
		flag := false
		var a, b string
		n, err := fmt.Scanf("%s %s\r\n", &a, &b)
		if n!=2||err != nil {
			break
		}
		strs1 := strings.Split(a, ".")
		strs2 := strings.Split(b, ".")

		l1 := len(strs1)
		l2 := len(strs2)

		//都有或者都没有
		if l1 == l2 {
			if l1 == 1 && a == b {
				flag = true
			} else if strs1[0]==strs2[0] && len(strings.Replace(strs2[1], "0", "", -1)) == len(strings.Replace(strs1[1], "0", "", -1)){
				flag=true
			}
		} else {
			if strs1[0] == strs2[0] {
				if len(strs1) == 1 && len(strings.Replace(strs2[1], "0", "", -1)) == 0 {
					flag = true
				}else if len(strs2) == 1 && len(strings.Replace(strs1[1], "0", "", -1)) == 0{
					flag=true
				}
			}
		}
		if flag {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
