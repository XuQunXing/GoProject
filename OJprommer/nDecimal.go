package OJprommer

import "fmt"

func main() {
	var a float64
	var n int
	for  {
		_,err:=fmt.Scanf("%f %d",&a,&n)
		if err !=nil {
			break
		}
		if a==0&&n== 0 {
			break
		}

		fmt.Print("0.")
		for i:=0;i<10 ;i++  {
			a  = a*(float64(n))
			var temp  = int(a)
			 fmt.Print(temp)
			a=a-float64(temp)
		}
		fmt.Println()
	}
}
