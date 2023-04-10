package main

import (
	"net/http"

	"github.com/andre-haniak/go-starter/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
