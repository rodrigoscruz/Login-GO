package account

import(
	"net/http"
	"encoding/json"

	user "localhost.com/controller/user"
)

func Index(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "./pages/account.html")
}

func ValidateData(w http.ResponseWriter, r *http.Request) {
	var data user.Data

	json.NewDecoder(r.Body).Decode(&data)

	user.NewUser(data)
}