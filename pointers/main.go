package main
import "fmt"
func main(){
	age :=44
	fmt.Println(age)
	x := &age //memory address
	fmt.Println(x)
	age = 20
	fmt.Println(x)
	fmt.Println(*x)
	fmt.Println(*x+35)
}

