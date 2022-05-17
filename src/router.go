package router

import (
	"net/http"

	account "localhost.com/account"
	login "localhost.com/login"
)

func HandleRoutes() {

	http.HandleFunc("/", login.Index)
	http.HandleFunc("/account", account.Index)
}
