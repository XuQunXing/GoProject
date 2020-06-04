package main

import (
	"fmt"
	"sort"
)

type County struct {
	name   string
	gold   int
	silver int
	bronze int
}
type Countwrapper struct {
	coun [] County
	by   func(p, q *County) bool
}

//重写len方法
func (pw Countwrapper) Len() int {
	return len(pw.coun)
}

//重写swap方法
func (pw Countwrapper) Swap(i, j int) {
	pw.coun[i], pw.coun[j] = pw.coun[j], pw.coun[i]
}

//重写less方法
func (pw Countwrapper) Less(i, j int) bool {
	return pw.by(&pw.coun[i], &pw.coun[j])
}
func main() {
	var n int
	var g, s, b int
	var c string
	fmt.Scanf("%d", &n)
	county := []County{}
	for i := 0; i < n; i++ {

		fmt.Scanf("%s %d %d %d", &c,&g,&s,&b )
		county=append(county,County{c, g, s, b})
	}
	//fmt.Println(county)

	sort.Sort(Countwrapper{county, func(p, q *County) bool {
		if p.gold != q.gold {
			return p.gold > q.gold
		} else if p.silver != q.silver {
			return p.silver > q.silver
		} else if p.bronze != q.bronze {
			return p.bronze > q.bronze
		}
		return p.name < q.name
	}})

	for _, value := range county {
		fmt.Println(value.name)
	}

}
