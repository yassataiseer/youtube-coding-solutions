
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)


func main(){
	db,err :=sql.Open("mysql","root:new_password@tcp(127.0.0.1)/mockweb")
	if err != nil{
		panic(err)
	}
	var query string
	query = "CREATE TABLE IF NOT EXISTS Users(Name VARCHAR(500),Password VARCHAR(500))"
	create,err := db.Exec(query)
	if err != nil{
		panic(err)
	}
	fmt.Println(create)


}