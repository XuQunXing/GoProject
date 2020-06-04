package main

import "fmt"

func main() {
	//map
	//一创建
	fmt.Println("----map的创建-----")
	//通过make函数创建
	dict:=make(map[string]int)
	fmt.Println(dict)
	//通过字面量创建
	map1:=map[string]string{"red":"#da1337","orange":"#e95a22"}
	fmt.Println(map1)
	//二、使用
	fmt.Println("----map的使用-----")
	colors:=map[string]string{}
	colors["red"]="#da1235"
	fmt.Println(colors)
	//不初始化map的话，创建的是nil map，它不能用来存放键值对

	// 判断map的键和值 第一种方法：  从map里面检索值，并判断键是否存在
	value,exists:=colors["red"]
	if exists {
		fmt.Println(value)
	}
	//第二种方法
	value1:=colors["red"]
	if value1!="" {
		fmt.Println(value1)
	}
	colors["sky"]="blue"
	fmt.Println(colors)
	//删除
	delete(colors,"sky")
	fmt.Println(colors)
}
