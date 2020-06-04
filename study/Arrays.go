package main

import (
	"fmt"
)

/**
集合学习
*/
func main() {
	//1、数组
	//声明一个长度为5的整型数组
	var arr [5]int
	fmt.Println(arr)
	//声明并初始化
	arr1:=[5]int{1,2,3,4,5}
	fmt.Println(arr1)
	//通过初始化值得个数来推导出数组容量
	arr2:=[...]int{1,2,3,4,5}
	fmt.Println(arr2)
	//元素来初始化
	arr3:=[5]int{1:1,2:2}
	fmt.Println(arr3)

	//指针数组
	arr4:=[5]*int{0:new(int),1:new(int)}
	*arr4[0]=8
	*arr4[1]=12
	fmt.Println(arr4)
	fmt.Println(*arr4[0],*arr4[1])

	//go中数组是一个值 所以可以使用=来进行赋值(前提类型和长度要一样)和输出
	var arr5 [5]int =arr3
	fmt.Println(arr5)

	//拷贝指针数组，实际上是拷贝指针的值，而不是指向的值
	var a [3]*string
	b := [3]*string{new(string), new(string), new(string)}
	*b[0] = "Red"
	*b[1] = "Blue"
	*b[2] = "Green"
	a = b
	fmt.Println(a,b)

	fmt.Println("------多维数组--------")
	//多维数组
	var array [4][2]int
	fmt.Println(array)
	array1:=[4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println(array1)
}
