package main

import (
	"net/http"

	"example.com/m/routes"
)

func main() {

	routes.LoadRoutes()

	// This will make the code listen in the 8000 port
	http.ListenAndServe(":8000", nil)
}
