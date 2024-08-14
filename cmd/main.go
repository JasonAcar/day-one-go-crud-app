package main

import (
	"github.com/JasonAcar/test-crud-app/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.StoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:42069", r))
}
