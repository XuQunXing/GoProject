package OJprommer

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	for true {
		var a int
		_, err := fmt.Scanf("%d", &a)
		if err != nil {
			break
		}
		standings:=standing(a)
		print(standings)
	}
}
//计算约数
func standing(a int) []int  {
	num:=float64(a)
	sqrt:=math.Sqrt(num)
	result:=make([]int,0,10)
	for i:=1;float64(i)<=sqrt ;i++  {
		if a%i==0 {
			i2:=a/i
			if i2==i {
				result=append(result,i)
				continue
			}
			result=append(append(result,i),i2)
		}
	}
	sort.Ints(result)
	return result
}
//打印输出
func print(arr[] int) {
	fmt.Print(len(arr))
	for _,value:=range arr {
		fmt.Printf(" %d",value)
	}
	fmt.Println()
}