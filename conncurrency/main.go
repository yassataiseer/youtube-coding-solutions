package main

import ("fmt"
)

func counter(n int) {
  for i := 0; i < n; i++ {
    fmt.Println(i)
  }
}

func main() {
	fmt.Println("HELLO THERE")
	go counter(200000)
	fmt.Scanln()
}