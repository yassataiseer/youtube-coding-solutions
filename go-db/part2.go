//Here is where we inject data we sql
//Here are the tables
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)


func main(){
	db,err :=sql.Open("mysql","root:new_password@tcp(127.0.0.1:3306)/golangdb")
	if err != nil{
		panic(err)
	}
	exec,err:= db.Query("INSERT INTO Users (Name,Password) VALUES(?,?)",("Clark Kent"),("Superman"))
	if err != nil{
		panic(err)
	}
	fmt.Println(exec)
	defer db.Close()
}
