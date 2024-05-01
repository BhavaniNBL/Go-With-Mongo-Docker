package main

import (
	router "ProjectGoMongo/Router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Mongo DB API")
	r := router.Router()
	fmt.Printf("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port: 4000")
}
