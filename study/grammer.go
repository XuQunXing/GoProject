package main

import "fmt"

func main() {
	var a,b,c int
	var err error
	fmt.Println(a,b,c,err)
	const pai  = 3.14
	//iota是自动递增枚举常量  初始为0 步进为1   注意：如果不在第一个引用，会按照组内常量数量递增。  仅在一个常量定义组里面有效，跳出归0
	const (
		h = "A"
		i = 'B'
		g = iota    // c = 2
		d = "D"
		e = iota    // e = 4
	)
	fmt.Println(g,e)

	//定义结构体
	type s struct {
		m int
		n string
	}
	//结构体的初始化  通过只作为列表来分配一个结构体
	var stru s  =s{
		m: 1,
		n: "test",
	}
	fmt.Println(stru.m,stru.n,stru)
}
