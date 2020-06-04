package main

import "fmt"

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
		//处理
		count:=0
		temp:=[]byte(bracket)
		var begin int
		//找到第一个左括号
		for index,value:=range temp{
			if value==')' {
				count++
				continue
			}else {
				begin=index
				break
			}
		}
		if count<len(temp) {
			cot:=0
			for i:=begin;i<len(temp);i++{
				if temp[i]=='(' {
					cot++
				}else {
					cot--
				}
				if cot<0 {
					count++
					cot=0
				}
			}
			if cot>0 {
				count+=cot
			}
		}


		fmt.Println(count)
	}
}
