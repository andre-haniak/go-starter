package main

import (
	"fmt"

	"github.com/andre-haniak/go-starter/models"
)

func main() {
	user := models.User{
		ID:        1,
		FirstName: "Andre",
		LastName:  "Haniak",
	}
	fmt.Println(user)
}
