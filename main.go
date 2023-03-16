package main

import (
	"log"
	routes "lucasdev-96/Routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	routes.HeroesRoutes(router)
	log.Fatal(http.ListenAndServe(":3000", router))

}
