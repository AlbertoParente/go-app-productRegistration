package controllers

import (
	"html/template"
	"net/http"

	"github.com/go-app-productRegistration/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()
	templates.ExecuteTemplate(w, "Index", allProducts)

}

func NewProduct(w, http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}
