package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func dbConnection() *sql.DB {

	// Setting the connections of the DataBase
	connection := "user=postgres dbname=store_project password=zeu$@2022 host=localhost sslmode=disable"

	// Open opens a database specified by its database driver name and a driver-specific data source name
	db, err := sql.Open("postgres", connection)

	// Catching the errors
	if err != nil {
		panic(err.Error())
	}

	return db
}

// The product struct
type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

// Taking the templates from the 'templates' folder with the Must function
var tmp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)
	// This will make the code listen in the 8000 port
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	db := dbConnection()

	allProducts, err := db.Query("select * from products")

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

		err = allProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	tmp.ExecuteTemplate(w, "Index", products)
	defer db.Close()
}
