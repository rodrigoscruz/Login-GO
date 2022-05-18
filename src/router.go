package router

import(
	"net/http"
	login "localhost.com/login"
	account "localhost.com/account"
)

func HandleRoutes() {

	http.HandleFunc("/", handleLogin)
	http.HandleFunc("/method", handleAccount)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	
	var methods = map[string] func() {
		"GET":    func() { login.Index(w, r)},
		"POST":   func() { login.ValidateLogin(w, r)},
	}

	validateRequest(methods, r.Method)
}

func handleAccount(w http.ResponseWriter, r *http.Request) {

	var methods = map[string] func() {
		"GET":    func() { account.Index(w, r)},
		"POST":   func() { account.ValidateData(w, r)},
	}

	validateRequest(methods, r.Method)
}

func validateRequest(methods map[string]func(), method string) {

	var function func()
	var found bool

	function, found = methods[method]

	if found {
		function()
	}
}