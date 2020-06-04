package main

import "fmt"

func main() {
	var n,m int
	for true {
		_,err:=fmt.Scanf("%d %d",&n,&m)
		if err!=nil {
			break
		}
		fmt.Print("+")
		for j:=0;j<m ;j++  {
			fmt.Print("---+")
		}
		fmt.Println()
		for i:=0;i<n;i++{

			fmt.Print("|")
			for j:=0;j<m ;j++  {
				fmt.Print("   |")
			}
			fmt.Println()

			fmt.Print("+")
			for j:=0;j<m ;j++  {
				fmt.Print("---+")
			}
			fmt.Println()
		}
	}
}
