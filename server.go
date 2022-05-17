package main

import (
	"net/http"
	"fmt"
	"localhost.com/router"
)


func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,`
	<h1>Olá mundo!</h1>
	<a href="/login">Login</a>
	`)
}	


func login(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "src/login.html")
}

func form(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "src/form.html")
}

func style(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "src/assets/login.css")
}

func stylef(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "src/assets/stylef.css")
}

func main() {

	router.HandleRoutes()

	http.HandleFunc("/", index)

	http.HandleFunc("/login", login)

	http.HandleFunc("/form", form)

	http.HandleFunc("/assets", style)

	http.HandleFunc("/assetsf", stylef)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8080", nil)
}
