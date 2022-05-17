package router

import(
	"net/http"

	login "localhost.com/login"
	accont "localhost.com/account"
)

func HandleRoutes() {

	http.HandleFunc("/", login.Index)
	http.HandleFunc("/account", account.Index)
}