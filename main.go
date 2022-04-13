package main

import (
	"html/template"
	"net/http"
)

// The product struct
type Product struct {
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

	// This slice will be set directly in the web server
	produtos := []Product{
		{Name: "Iphone 8s", Description: "64GB", Price: 450.00, Quantity: 1},
		{"PS4", "1T", 500.00, 5},
		{"XBOX ONE", "500GB", 350.00, 4},
	}

	tmp.ExecuteTemplate(w, "Index", produtos)
}
