package main

import (
	routes "lucasdev-96/Routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.HeroesRoutes(router)

}
