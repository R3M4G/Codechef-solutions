package main

import (
	"fmt"
)

func main() {
	var i int
	var j float64
	fmt.Scanln(&i,&j)
	if i % 5 == 0 && i < int(j) {
		fmt.Println((j-float64(i))-0.50)
	} else {
		fmt.Printf("%.2f",j)
	}
}

