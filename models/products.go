package models

import (
	"example.com/m/db"
)

// The product struct
type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func SearchAllProducts() []Product {
	db := db.DbConnection()

	allProducts, err := db.Query("select * from products order by id asc")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	// With the function Next, all the columns and lines from the query will be checked
	for allProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		// Scan copies the columns in the current row into the values pointed at by dest.
		// The number of values in dest must be the same as the number of columns in Rows.
		err = allProducts.Scan(&id, &name, &description, &price, &quantity)
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

	// CLosing the database
	defer db.Close()

	return products
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := db.DbConnection()

	dbInsertion, err := db.Prepare("insert into products(nome, description, price, quantity) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	dbInsertion.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(id string) {

	db := db.DbConnection()
	deleteProduct, err := db.Prepare("delete from products where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Product {

	db := db.DbConnection()

	dbProduct, err := db.Query("select * from products where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Product{}

	for dbProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = dbProduct.Scan(&id, &name, &description, &price, &quantity)
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

func UpdateProduct(id int, name, description string, price float64, quantity int) {

	db := db.DbConnection()

	updateProduct, err := db.Prepare("update products set nome=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	updateProduct.Exec(name, description, price, quantity, id)

	defer db.Close()

}
