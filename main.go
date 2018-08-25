package main

import (
	"medhackmty/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/sintomas", controllers.GetInit)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
