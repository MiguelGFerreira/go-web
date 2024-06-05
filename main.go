package main

import (
	"net/http"
	"pkg/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
