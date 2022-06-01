package main

import (
	"fmt"
	"log"
	"net/http"

	Routes "github.com/aman98072/crudAPI/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	Routes.RegisterUserRoutes(r)
	http.Handle("/", r)
	fmt.Println("Listing for requests at http://localhost:9010/")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
