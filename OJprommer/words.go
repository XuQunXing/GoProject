package OJprommer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scan:=bufio.NewScanner(os.Stdin)
	scan.Scan()
	line:=scan.Text()
	fmt.Printf("%.2f\n",float64(len(line)-len(strings.Split(line," "))+1)/float64(len(strings.Split(line," "))))
}
