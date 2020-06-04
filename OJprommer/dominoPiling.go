package main

import (
	"fmt"
)

/*
48. Domino piling
 */

func main() {
	var m,n int
	for true {
		_,err:=fmt.Scanf("%d %d",&m,&n)
		if err!=nil {
			break
		}
		if m==1 {
			fmt.Println(n/2)
			continue
		}
		if n==1 {
			fmt.Println(m/2)
			continue
		}
		//横排
		heng:=m*(n/2)
		if n%2!=0 {
			heng+=m/2
		}
		//竖排
		shu:=n*(m/2)
		if m%2!=0 {
			shu+=n/2
		}
		if heng>=shu {
			fmt.Println(heng)
		}else {
			fmt.Println(shu)
		}
	}
}
