package router

import (
	"net/http"
	"html/template"
)

func Index(w http.ResponseWriter, r *http.Request){
	template.Must(template.ParseFiles("Views/Index.html"))
}

func Register(w http.ResponseWriter, r *http.Request){
	template.Must(template.ParseFiles("Views/Registe.html"))
}

func Login(w http.ResponseWriter, r *http.Request){
	template.Must(template.ParseFiles("Views/Login.html"))
}

func Logout(w http.ResponseWriter, r *http.Request){
	template.Must(template.ParseFiles("Views/Logout.html"))
}