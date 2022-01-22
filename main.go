package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	connection := "user=postgres dbname=product password=************ host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Product struct {
	id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := ConnectDatabase()

	getAllProducts, err := db.Query("select * from product")
	if err != nil {
		panic(err.Error())
	}

	p := Product()
	products := []Product{}

	for getAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = getAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	templates.ExecuteTemplate(w, "Index", products)
	defer db.Close()
}
