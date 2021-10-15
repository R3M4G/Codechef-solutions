package main
import "fmt"

func main(){
	var L,B int
	fmt.Scan(&L)
	fmt.Scan(&B)
	if L*B > 2*(L+B) {
		fmt.Println("Area")
		fmt.Println(L*B)
	 } else if L*B == 2*(L+B) {
        fmt.Println("Eq")
        fmt.Println(L*B)
	} else {
		fmt.Println("Peri")
		fmt.Println(2*(L+B))
	}
}