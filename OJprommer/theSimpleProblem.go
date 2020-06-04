package main

import "fmt"

/*
24
 */
func main() {
	var a,k,n,total int
	for true {
		_,err:=fmt.Scanf("%d %d %d",&a,&k,&n)
		if err!=nil {
			break
		}
		total=(n+1)*a+k*(1+n)*n/2
		fmt.Println(total)
	}
}
