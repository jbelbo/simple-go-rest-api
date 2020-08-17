package main

import (
	"net/http"

	"github.com/jbelbo/simple-go-rest-api/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
