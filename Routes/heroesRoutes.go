package routes

import (
	controller "lucasdev-96/Controller"

	"github.com/gorilla/mux"
)

func HeroesRoutes(router *mux.Router) {
	router.HandleFunc("/heroes", controller.GetHeroes).Methods("GET").Name("heroes")
	router.HandleFunc("/heroes/{name}", controller.GetHero).Methods("GET").Name("hero")
}
