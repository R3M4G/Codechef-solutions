package main
import "fmt"

func main(){
	// your code goes here
	var t,i,j int
	fmt.Scanln(&t)
	for l := 1; l <= t; l++ {
	fmt.Scanln(&i,&j)
	fmt.Println(i/j)	
	}
}
