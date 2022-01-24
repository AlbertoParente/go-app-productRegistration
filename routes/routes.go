package routes

import (
	"net/http"

	"github.com/go-app-productRegistration/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
