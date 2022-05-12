package main

import (
	"net/http"
	"fmt"
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

func style(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "src/assets/style.css")
}


func main() {
	http.HandleFunc("/", index)

	http.HandleFunc("/login", login)

	http.HandleFunc("/assets", style)

	http.ListenAndServe(":80", nil)
}
