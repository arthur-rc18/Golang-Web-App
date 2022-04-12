package main

import (
	"net/http"
)

func main() {

	// This will make the code listen in the 8000 port
	http.ListenAndServe(":8000", nil)
}
