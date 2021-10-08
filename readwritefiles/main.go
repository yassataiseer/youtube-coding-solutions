package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func write_File(){
	file, err := os.Create("test.txt")
	if(err!=nil){
		panic(err)
	}
	defer file.Close()
	_,err = file.WriteString("Hello world from Go")
	if(err!=nil){
		panic(err)
	}
}
func read_File(){
	data,err := ioutil.ReadFile("test.txt")
	if(err!=nil){
		panic(err)
	}
	fmt.Println(string(data))
}
func main(){
	write_File()
	read_File()
}