package main
import "fmt"
func main(){
    var first int
    fmt.Println("Enter Your First Number: ")
    fmt.Scanln(&first)
    var scnd int
    fmt.Println("Enter Your Secound Number: ")
    fmt.Scanln(&scnd)
	fmt.Println("Your answer is:")
	fmt.Println(add(first,scnd));
}
func add( x int, y int) int{
	return x+y
}

