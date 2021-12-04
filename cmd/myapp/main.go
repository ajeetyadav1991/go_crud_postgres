package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ajeetyadav1991/go_postgres_crud/internal"
)

func main() {

	//databases.CreateConnection() //optional

	// The main.go is our server will serve all the Router using the
	// Router function created in routes.go in internal package
	r := internal.Router()

	fmt.Println("Server Started...")
	log.Fatal(http.ListenAndServe(":4000", r))
}
