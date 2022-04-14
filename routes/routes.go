package routes

import (
	"net/http"

	"example.com/m/controller"
)

func LoadRoutes() {

	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.New)
	http.HandleFunc("/insert", controller.Insert)
}
