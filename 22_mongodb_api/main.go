package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bikashsaud/mongodb_api/router"
)

func main() {
	fmt.Println("Create MongoDB API Setup")
	fmt.Println("Starting Server.")
	r := router.Router()

	fmt.Println("Server started...")

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listing at port 4000")

}
