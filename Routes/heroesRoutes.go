package routes

import (
	controller "lucasdev-96/Controller"

	"github.com/gorilla/mux"
)

func HeroesRoutes(router *mux.Router) {
	getHeroes := controller.GetHeroes
	router.HandleFunc("/heroes", getHeroes)
}
