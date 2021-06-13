package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)
func main(){
    files,err := ioutil.ReadDir("./")
    if err != nil{
        panic(err)
    }
    for _, f:=range files{
        fmt.Println(f.Name())
        if(f.Name()=="main.go"){

        }else{
            start(f.Name())
        }
    }
}
func start(filename string){
    path := filename
    removeLine(path,0)
}


func removeLine(path string, lineNumber int){
    file,err := ioutil.ReadFile(path)
    if err!=nil{
        panic(err)
    }
    info,_:= os.Stat(path)
    mode := info.Mode()
    array:= strings.Split(string(file),"\n")
    array = append(array[:lineNumber],array[lineNumber+1:]...)
    ioutil.WriteFile(path,[]byte(strings.Join(array,"\n")),mode)
}

