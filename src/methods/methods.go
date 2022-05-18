package methods

import (
	"net/http"      // Implementa cliente/servidor HTTP.
	"encoding/json" // Codificação em JSON.
	"strings"       // Manipula strings codificadas em UTF-8.
)

var data struct {

	Method string `json: method`

	Descrition string `json: descrition`
}

func Index(w http.ResponseWriter, r *http.Request) {

	data.Method = r.Method
	data.Descrition = " - Apresentar recursos."
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func Create(w http.ResponseWriter, r *http.Request) {

	data.Method = r.Method
	data.Descrition = " - Recurso criado."
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func Update(w http.ResponseWriter, r *http.Request) {

	data.Method = r.Method
	data.Descrition = " - Recurso atualizado."
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNoContent)
}

func Search(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	var name = params.Get("name")
	var date = strings.ReplaceAll(params.Get("date"), "-", "/")
	data.Method = r.Method
	data.Descrition = " - Buscar por: " + name + "e data - " + date
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func Show(w http.ResponseWriter, r *http.Request) {

	var parts []string = strings.Split(r.URL.Path, "/")
	data.Method = r.Method
	data.Descrition = " - Apresentar recurso com ID - " + parts[3]
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}