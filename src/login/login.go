package login

import(
	"net/http"
	"fmt"

	user "localhost.com/controller/user"
)

func Index(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "./pages/login.html")
}

func ValidateLogin(w http.ResponseWriter, r *http.Request) {

	var email = r.PostFormValue("email")
	var password = r.PostFormValue("password")

	if user.SearchUser(email, password) {

		fmt.Fprintf(w,`
			<h1>Iniciar novo acesso</h1>
			<p>Método: %s</p>
			<p>Email: %s</p>
			<p>Senha: %s</p>`,
			r.Method, email, password)

		return
	}

	http.Redirect(w, r, "./login", 301)
}