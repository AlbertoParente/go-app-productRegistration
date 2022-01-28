package routes

import (
	"net/http"

	"github.com/go-app-productRegistration/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.NewProduct)
	http.HandleFunc("/insert", controllers.Insert)
}
