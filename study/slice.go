package main

import "fmt"

func main() {
	//切片
	//一、创建
	//第一次中创建方法：使用内建函数make
	slice:=make([]string,5)
	fmt.Println(slice)
	//声明长度和容量
	slice1:=make([]string,3,5)
	fmt.Println(slice1)
	//第二种创建方法：使用slice字面量
	slice2:=[]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Println(slice2)

	//初始化索引    初始化一个有100个元素的空的字符串
	slice3 := []string{99: ""}
	fmt.Println(slice3)

	//创建一个nil slice的方法，声明但部署池化
	var slice4 []int
	fmt.Println(slice4)
	//创建一个empty slice  声明并初始化一下
	slice5 :=make([]int,0)
	slice6:=[]int{}
	fmt.Println(slice5,slice6)

	//二、使用
	fmt.Println("slice使用------------")
	//赋值与数组一样
	//内建函数  append,len,cap
	sli := []int{10, 20, 30, 40, 50}
	// 长度为2，容量为4
	newSlice := sli[1:3]
	fmt.Println(newSlice)
	/*
	建立在同一个上面的slice， 共享其中的数据,既同时改变。（注意索引要对应上）
	对于 slice[i:j] 和底层容量为 k 的数组
	长度：j - i
	容量：k - i
	 */
	sli[1]=99
	fmt.Println(sli[1],newSlice[0])

	//三、增长
	fmt.Println("slice增长------------")
	//如果有底层数组，并且有可用容量的话，直接修改的是底层数组。
	//大于可以用容量的话，在1000以内的话每次扩充是两倍，大于1000扩充1.25倍。
	newSlice=append(newSlice,60)
	haha:=append(sli,166)
	fmt.Println(newSlice,haha)
	//限定容量增长，不会影响原来底层数组
	source := []string{"apple", "orange", "plum", "banana", "grape"}
	temp := source[2:3:3]
	temp = append(temp, "kiwi")
	fmt.Println(source,temp)

	//四、迭代
	fmt.Println("slice迭代------------")
	//&slice[index]  可以得到真正的地址
	//如果需要索引值的话可以加上index，不需要的话可以用_来代替
	die := []int{10, 20, 30, 40, 50}
	for index,value:=range die {
		fmt.Printf("Index:%d Vaule:%d\n",index,value)
	}
}
