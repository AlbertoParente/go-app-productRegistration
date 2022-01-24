package main

import (
	"net/http"

	"github.com/go-app-productRegistration/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
