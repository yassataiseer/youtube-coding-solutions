
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
	//This is the struct that will structure db data


func main(){
	db,err :=sql.Open("mysql","YOUR_NAME:YOUR_PSSWD@tcp(YOUR_IP_ADDRESS)/YOUR DB NAME")
	if err != nil{
		panic(err)
	}
	rows,err:= db.Query("SELECT * FROM USERS")//Grab data from db
	if err != nil{
		panic(err)
	}
	var Username string
	var Password string
	//Create variables 
	var data []User_data
	// Make an struct instance of User_data
	for rows.Next(){
		err = rows.Scan(&Username,&Password)//Scan rows
		if err != nil{
			panic(err)
		}
		data = append(data, User_data{Username:Username,Password: Password})//append to data stuct
	}
	fmt.Println(data)
	fmt.Println(data[0].Username)
	defer db.Close()
}
