package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc:=bufio.NewScanner(os.Stdin)
	for  {
		var n int
		_,err:=fmt.Scanln(&n)
		if err!=nil {
			break
		}
		for i:=1;i<=n ;i++  {
			sc.Scan()
			temp:=strings.Split(sc.Text()," ")
			var a,b int
			a, _ =strconv.Atoi(temp[0])
			b, _ =strconv.Atoi(temp[1])
			addNum:=a+b
			fmt.Printf("Case %d:\n",&i)
			fmt.Printf("%d + %d = %d\n",&a,&b,&addNum)
		}
	}
}
