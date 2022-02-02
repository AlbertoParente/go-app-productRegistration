package models

import (
	"github.com/go-app-productRegistration/db"
	"github.com/go-app-productRegistration/models"
)

type Product struct {
	id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {
	db := db.ConnectDatabase()

	getAllProducts, err := db.Query("select * from product")
	if err != nil {
		panic(err.Error())
	}

	p := models.Product()
	products := []Product{}

	for getAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = getAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := db.ConnectDatabase()

	insertDataIntoTheDatabase, err := db.Prepare("insert into produtos(name, description, price, quantity) values($1, $2. $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertDataIntoTheDatabase.Exec(name, description, price, quantity)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectDatabase()
	deleteByProduct, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteByProduct.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.ConnectDatabase()
	productDatabase, err := db.Query("select * from product where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Product{}

	for productDatabase.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productDatabase.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		productToUpdate.Id = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Quantity = quantity
	}
	defer db.Close()
	return productToUpdate
}

func UpdateProduct(id int, name, description string, price, float64, quantity int) {
	db := db.ConnectDatabase()

	UpdateProduct, err := db.Prepare("update product set name=$1, description=$2, price=$3 quantiry=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	UpdateProduct.Exec(name, description, price, quantity, id)
	defer db.Close()
}
