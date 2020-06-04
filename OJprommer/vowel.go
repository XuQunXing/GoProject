package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc:=bufio.NewScanner(os.Stdin)
	sc.Scan()
	line:=sc.Text()
	//vowel:=map[string]string{"a":"A","e":"E","i":"I","o":"O","u":"U"}
	//consonant:=map[string]string{"B":"b","C":"c","D":"d","F":"f","G":"g","H":"h","J":"j","K":"k","L":"l","M":"m","N":"n","P":"p","Q":"q","R":"r","S":"s","T":"t","V":"v","W":"w","X":"x","Y":"y","Z":"z"}
	temp:=strings.ToLower(line)
	//println(temp)
	r:=[]rune(temp)
	for i:=0;i<len(r);i++{
		if r[i]==97|| r[i]==101|| r[i]==105|| r[i]==111|| r[i]==117{
			r[i]=r[i]-32
		}
	}
	fmt.Println(string(r))
}
