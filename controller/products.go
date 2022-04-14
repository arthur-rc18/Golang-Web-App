package controller

import (
	"html/template"
	"net/http"
	"strconv"

	"example.com/m/models"
)

// Taking the templates from the 'templates' folder with the Must function
var tmp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	allProducts := models.SearchAllProducts()
	tmp.ExecuteTemplate(w, "Index", allProducts)

}

func New(w http.ResponseWriter, r *http.Request) {
	tmp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	// Method specifies the HTTP method (GET, POST, PUT, etc.). For client requests, an empty string means GET
	if r.Method == "POST" {
		// FormValue returns the first value for the named component of the query
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		// Converting the data in string to float64
		convertedPriceF, err := strconv.ParseFloat(price, 64)
		if err != nil {
			panic(err.Error())
		}

		// Converting the data in string to int
		convertedQuantityI, err := strconv.Atoi(quantity)
		if err != nil {
			panic(err.Error())
		}

		models.CreateNewProduct(name, description, convertedPriceF, convertedQuantityI)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.EditProduct(productId)
	tmp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idConverted, err := strconv.Atoi(id)
		if err != nil {
			panic(err.Error())
		}

		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			panic(err.Error())
		}

		quantityConverted, err := strconv.Atoi(quantity)
		if err != nil {
			panic(err.Error())
		}

		models.UpdateProduct(idConverted, name, description, priceConverted, quantityConverted)
	}

	http.Redirect(w, r, "/", 301)
}
