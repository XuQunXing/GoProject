package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc:=bufio.NewScanner(os.Stdin)
	sc.Scan()
	line:=sc.Text()
	strs:=strings.Split(line," ")
	var ans string
	for _,value:=range strs{
		for i:=len(value)-1;i>=0 ;i-- {
			ans+=string(value[i])
		}
		ans+=" "
	}

	fmt.Println(strings.TrimSpace(ans))
}
