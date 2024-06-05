package routes

import (
	"net/http"
	"pkg/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
