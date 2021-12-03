package routes

import (
	"fmt"
	"net/http"
	"projeto-go/controllers"
	_ "projeto-go/controllers"

)

func LoadRoutes(){
	fmt.Println("Server up on port 8000 ")
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
}
