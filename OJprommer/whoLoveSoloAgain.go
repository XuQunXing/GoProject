package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"unicode"
)

/*
26. Who Love Solo Again
 */

func main() {
	input := bufio.NewReader(os.Stdin)
	for  {
		line, _ := input.ReadString('\n')
		if len(line) <= 1 {
			return
		}
		word:=words(line)
		var ans string
		if len(word)==0 {
			ans= "."
		}else {
			for i:=0;i<len(word);i++{
				if i!=len(word)-1 {
					ans+=upper(word[i])+" "
				}else {
					ans+=upper(word[i])+"."
				}

			}
		}
		fmt.Println(ans)
	}
}
//单词首字母大写
func upper(s string) string{
	temp:=[]byte{s[0]}
	tempers :=string(bytes.ToUpper(temp))
	return tempers +s[1:len(s)]
}
//提取单词
func words(s string) []string {
	i:=0
	ans:=make([]string,100)
	count:=0
	sBy:=[]rune(s)
	for i<len(sBy)  {
		if !unicode.IsLetter(sBy[i]) {
			i++
			continue
		}else {
			var temp string=""
			temp+=string(sBy[i])
			i++
			for unicode.IsLetter(sBy[i]) {
				temp+=string(sBy[i])
				i++
			}
			ans[count]=temp
			count++
		}
	}
	total:=ans[0:count]
	return total
}