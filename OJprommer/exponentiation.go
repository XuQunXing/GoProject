package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

/*
20. Exponentiation
*/
func main() {
	input := bufio.NewReader(os.Stdin)
	for {
		line, _ := input.ReadString('\n')
		if len(line) <= 1 {
			return
		}
		line = strings.TrimSpace(line)
		numarr := strings.Split(line, " ")
		n, _ := strconv.Atoi(numarr[1])
		temp := strings.Split(numarr[0], ".")
		pointlength := len(temp[1])

		//数字转化为整数
		num, _ := strconv.ParseFloat(numarr[0], 64)
		intR := int64(num * math.Pow10(pointlength))
		R := big.NewInt(intR)
		ans := big.NewInt(1)
		for i := 0; i < n; i++ {
			ans.Mul(ans, R)
		}
		//得到总的小数位数，整数结果转化为字符串处理
		finalPointlength := pointlength * n
		ansStr := ans.String()

		//判断小数点在哪
		if len(ansStr) > finalPointlength {
			ansStr = ansStr[0:len(ansStr)-finalPointlength] + "." + ansStr[len(ansStr)-finalPointlength:len(ansStr)]
		} else if len(ansStr) == finalPointlength {
			ansStr = "." + ansStr
		} else {
			zeroStr := "."
			for i := 0; i < finalPointlength-len(ansStr); i++ {
				zeroStr += "0"
			}
			ansStr = zeroStr + ansStr
		}
		//去除后面的零
		for ansStr[len(ansStr)-1] == []byte("0")[0] {
			ansStr = ansStr[0 : len(ansStr)-1]
		}
		fmt.Println(ansStr)
	}

}
