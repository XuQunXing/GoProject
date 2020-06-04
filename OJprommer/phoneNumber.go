package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {
	r := bufio.NewReader(os.Stdin)

	var parameters [][]string
	for {
		input, _ := r.ReadString('\n')
		input = strings.TrimSpace(input)
		if string(input)=="0" || len(input)==0 {
			break
		}
		n, _ := strconv.Atoi(input)
		var tempParameters []string
		for len(tempParameters)<n {
			newline, _ := r.ReadString('\n')
			newline = strings.TrimSpace(newline)
			tempParameters=append(tempParameters,newline)
		}
		parameters=append(parameters,tempParameters)
	}
	for _,phonenumList := range parameters {
		flag := true
		for i:=0;i<len(phonenumList);i++ {
			for j:=i+1;j<len(phonenumList);j++ {
				if len(phonenumList[j]) > len(phonenumList[i]) {
					if strings.HasPrefix(phonenumList[j],phonenumList[i]) {
						fmt.Println("NO")
						flag=false
						break
					}
				}else {
					if strings.HasPrefix(phonenumList[i],phonenumList[j]) {
						fmt.Println("NO")
						flag=false
						break
					}
				}
			}
			if flag==false {
				break
			}
		}
		if flag==true {
			fmt.Println("YES")
		}
	}
}