package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/user/api/model"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to homepage!")
	fmt.Println("Loading homepage!!")
}

func setheaders(w http.ResponseWriter, statusCode int, headers map[string]string) http.ResponseWriter {

	for headerKey, headerValue := range headers {
		w.Header().Set(headerKey, headerValue)
	}
	w.WriteHeader(statusCode)
	return w
}

func api(w http.ResponseWriter, r *http.Request) {

	headers := make(map[string]string)

	headers["Server"] = "A Go Web Server"
	headers["Content-type"] = "application/json"

	var statusCode int = 200

	w = setheaders(w, statusCode, headers)
	json.NewEncoder(w).Encode(model.LoadArticles())
	fmt.Println("Loading apis!!")
}

// LoadRoutes will load all apis used in this app
func LoadRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/api", api)
}
