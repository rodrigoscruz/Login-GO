package main

import (
	"net/http"
	"localhost.com/router"
)

func main() {

	router.HandleRoutes()

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8080", nil)
}
