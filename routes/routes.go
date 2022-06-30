package routes

import (
	"net/http"
	"github.com/biraneves/alr-go-web/controllers"
)

func LoadRoutes() {

	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)

}