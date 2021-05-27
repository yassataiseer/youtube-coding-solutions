
//We will fetch data username&psswd from db from the Users table
// Then structure it via structs

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)
type User_data struct {
	
	Username,Password string
	}


func main(){
	db,err :=sql.Open("mysql","root:new_password@tcp(127.0.0.1:3306)/golangdb")
	if err != nil{
		panic(err)
	}
	rows,err:= db.Query("SELECT * FROM USERS")
	if err != nil{
		panic(err)
	}
	var Username string
	var Password string
	var data []User_data
	for rows.Next(){
		err = rows.Scan(&Username,&Password)
		if err != nil{
			panic(err)
		}
		data = append(data, User_data{Username:Username,Password: Password})
	}
	fmt.Println(data)
	defer db.Close()
}
