package main

import "fmt"

func main() {
	var a int
	for true {
		_,err:=fmt.Scanf("%d",&a)
		if err !=nil{
			break
		}
		if (a&(a-1))==0  {
			fmt.Println("NO")
		} else{
			fmt.Println("YES")
		}
	}
}
