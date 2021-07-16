package main

import (
	"fmt"
	"html/template"
	"net/http"
)


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

	var tmplt = template.Must(template.ParseFiles("templates/index.html"))
	tmplt.Execute(w,nil)
}
func login_user(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	var Username = r.Form["fname"]
	var Password = r.Form["Password"]
	fmt.Println(Username," ",Password)

	var tmplt = template.Must(template.ParseFiles("templates/index.html"))
	tmplt.Execute(w,nil)
}

func main(){

	http.HandleFunc("/",login)
	http.HandleFunc("/sign_up",sign_up)
	http.HandleFunc("/login_user",login_user)
	http.HandleFunc("/signup_user",signup_user)
	http.ListenAndServe(":8000",nil)
}

