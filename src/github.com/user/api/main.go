package main

import (
	"log"
	"net/http"

	"github.com/user/api/routes"
)

func startServer() {
	routes.LoadRoutes()
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
	startServer()
}
