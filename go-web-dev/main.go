package main
import(
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request){
	var tplt = template.Must(template.ParseFiles("templates/login.html"))
	tplt.Execute(w,nil)	
}

func sign_up(w http.ResponseWriter, r *http.Request){
	var tplt = template.Must(template.ParseFiles("templates/sign_up.html"))
	tplt.Execute(w,nil)	
}

func main(){
	//Main function simply handles routing
	http.HandleFunc("/", login) 
	http.HandleFunc("/sign_up", sign_up)

	http.ListenAndServe(":8000",nil)
}