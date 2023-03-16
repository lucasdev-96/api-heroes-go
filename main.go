package main

import (
	"log"
	models "lucasdev-96/Models"
	routes "lucasdev-96/Routes"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	models.CreateHeroes()
}

func main() {
	router := mux.NewRouter()

	routes.HeroesRoutes(router)
	log.Fatal(http.ListenAndServe(":3000", router))

}
