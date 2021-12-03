package main

import (
	"net/http"
	_ "projeto-go/models"
	"projeto-go/routes"
)

func main() {
	routes.LoadRoutes();
	http.ListenAndServe(":8000", nil)
}

