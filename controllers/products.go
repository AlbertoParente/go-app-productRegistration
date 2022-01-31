package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/go-app-productRegistration/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()
	templates.ExecuteTemplate(w, "Index", allProducts)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedPriceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error in price conversion:", err)
		}

		convertedQuantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Quantity conversion error!:", err)
		}

		models.CreateNewProduct(name, description, convertedPriceFloat, convertedQuantityInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	models.DeleteProduct(idProduct)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	template.ExecuteTemplate(w, "Edit", nil)
}
