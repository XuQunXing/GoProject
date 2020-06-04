package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
39. 24点游戏
*/

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for {
		//输入数据
		if sc.Scan(){
			temp:=strings.Split(sc.Text()," ")
			//转化
			poke:=convert(temp)
			//处理
			if jugdeTwentyFour(poke, 4) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}else {
			break
		}
	}
}
//字符串转化为float
func convert(temp []string) []float64  {
	poke:=make([]float64,4)
	for index, value := range temp {
		switch value {
		case "A":
			poke[index] = float64(1)
		case "T":
			poke[index] = float64(10)
		case "J":
			poke[index] = float64(11)
		case "Q":
			poke[index] = float64(12)
		case "K":
			poke[index] = float64(13)
		default:
			poke[index], _ = strconv.ParseFloat(value, 64)
		}
	}
	return poke
}
//判断是否能组成24点
func jugdeTwentyFour(poke []float64, n int) bool {
	if n == 1 {
		if math.Abs(poke[0]-24) == 0 {
			return true
		} else {
			return false
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a := poke[i]
			b := poke[j]
			poke[j] = poke[n-1]
			//穷举两个数的可能性       减法和除法要分别进行两次
			poke[i] = a + b
			if jugdeTwentyFour(poke, n-1) {
				return true
			}
			poke[i] = a - b
			if jugdeTwentyFour(poke, n-1) {
				return true
			}
			poke[i] = b-a
			if jugdeTwentyFour(poke, n-1) {
				return true
			}
			poke[i] = a * b
			if jugdeTwentyFour(poke, n-1) {
				return true
			}
			if b != 0 {
				poke[i] = a / b
			}
			if jugdeTwentyFour(poke, n-1) {
				return true
			}
			if a != 0 {
				poke[i] = b / a
			}
			if jugdeTwentyFour(poke, n-1) {
				return true
			}
			//回溯
			poke[i] = a
			poke[j] = b
		}
	}
	return false
}
