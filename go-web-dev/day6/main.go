package main

import (
	"fmt"
	"database/sql"
	"html/template"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)


type user_data struct{
	Username string
	Password string
}

func grab_users() []user_data{
	db,err := sql.Open("mysql","root:new_password@tcp(127.0.0.1)/mockweb")
	if err!= nil{
		panic(err)
	}
	rows,err := db.Query("SELECT * FROM Users")
	if(err!=nil){
		panic(err)
	}	
	var username string
	var password string
	var user_rows []user_data
	for rows.Next(){
		err := rows.Scan(&username,&password)
		if(err!=nil){
			panic(err)
		}
		user_rows = append(user_rows, user_data{Username: username, Password: password})
	}
	defer db.Close()
	return user_rows
}




func add_user(Username string, password string)bool{
	db,err := sql.Open("mysql","root:new_password@tcp(127.0.0.1)/mockweb")
	if err!= nil{
		panic(err)
	}
	add,err := db.Query("INSERT INTO Users(Name,Password) VALUES (?,?)",(Username),(password))
	if err!= nil{
		panic(err)
	}
	fmt.Println(add)
	defer db.Close()
	return true
}

func check_user(Username string, password string)bool{
	db,err := sql.Open("mysql","root:new_password@tcp(127.0.0.1)/mockweb")
	if err!= nil{
		panic(err)
	}
	var exists bool
	var query string
	query = fmt.Sprintf("SELECT EXISTS(SELECT Name FROM Users WHERE Name='%s' AND Password='%s')",(Username),(password))
	row := db.QueryRow(query).Scan(&exists)
	if err!=nil && err!= sql.ErrNoRows{
		panic(err)
	}
	fmt.Println(row)
	defer db.Close()
	return exists
}
func login(w http.ResponseWriter, r *http.Request){
	var tmplt = template.Must(template.ParseFiles("templates/login.html"))
	tmplt.Execute(w,nil)
}
func sign_up(w http.ResponseWriter, r *http.Request){
	var tmplt = template.Must(template.ParseFiles("templates/sign_up.html"))
	tmplt.Execute(w,nil)
}
func signup_user(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	var Username = r.Form["fname"]
	var Password = r.Form["Password"]
	fmt.Println(Username," ",Password)
	if(add_user(Username[0],Password[0])){
		var tmplt = template.Must(template.ParseFiles("templates/index.html"))
		user := http.Cookie{
			Name: "Username", Value: Username[0],
		}
		http.SetCookie(w, &user)
		var cookie = user.Value
		fmt.Println("COOKIE: "+cookie)
		var user_data = grab_users()
		tmplt.Execute(w,user_data)
	}else{
		var tmplt = template.Must(template.ParseFiles("templates/error.html"))
		tmplt.Execute(w,nil)
	}

}
func login_user(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	var Username = r.Form["fname"]
	var Password = r.Form["Password"]
	fmt.Println(Username," ",Password)
	if(check_user(Username[0],Password[0])){
		var tmplt = template.Must(template.ParseFiles("templates/index.html"))
		user := http.Cookie{
			Name: "Username", Value: Username[0],
		}
		http.SetCookie(w, &user)
		var cookie = user.Value
		fmt.Println("COOKIE: "+cookie)
		var user_data = grab_users()

		tmplt.Execute(w,user_data) 
	}else{
		var tmplt = template.Must(template.ParseFiles("templates/error.html"))
		tmplt.Execute(w,nil)
	}
}

func main(){

	http.HandleFunc("/",login)
	http.HandleFunc("/sign_up",sign_up)
	http.HandleFunc("/login_user",login_user)
	http.HandleFunc("/signup_user",signup_user)
	http.ListenAndServe(":8000",nil)
}

