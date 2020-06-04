package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	fmt.Println("开始Goroutines")
	wg.Add(1)

	go func() {
		defer wg.Done()

		count := 0
		var flag bool
		for count < 3 {
			//TODO 发接口，判断返回数据，如果为success，flag置为true，break。fail的话，先等待300秒，count++。
			fmt.Println("第",count,"次订阅,结果",getData())
			if getData() == true {
				flag = true
				break
			} else {
				time.Sleep(time.Second*6)
				count++
			}
		}
		if flag==true {
			//TODO 返回对应的信息
			fmt.Println("订阅成功")
		}else {
			//TODO 打印相关的信息
			fmt.Println("订阅失败")
		}
	}()

	fmt.Println("Waiting To finish")
	wg.Wait()
}

//模拟返回数据
func getData() bool {
	return false
}
