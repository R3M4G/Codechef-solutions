package main
import (
	"fmt"
)

func main() {
	var i int
	var j float64
	fmt.Scan(&i,&j)
	if i % 5 != 0 && j < (float64(i)+0.5) {
		fmt.Printf("%.2f\n",j)		
	} else {
		fmt.Printf("%.2f\n", j-(float64(i)+0.5))
	}
}
