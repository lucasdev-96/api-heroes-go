package controller

import (
	"encoding/json"
	models "lucasdev-96/Models"
	"net/http"

	"github.com/gorilla/mux"
)

type Err struct {
	Error string `json: "error"`
}

func GetHeroes(response http.ResponseWriter, request *http.Request) {
	heroes := models.AllHeroes()
	json.NewEncoder(response).Encode(heroes)
}

func GetHero(response http.ResponseWriter, request *http.Request) {
	value := mux.Vars(request)
	hero, err := models.GetHero(value["name"])
	if err != nil {
		textErr := Err{Error: err.Error()}
		json.NewEncoder(response).Encode(textErr)
		return
	}
	json.NewEncoder(response).Encode(hero)
}

// func AddNewHero(response http.ResponseWriter, request *http.Request) {
// 	data := mux.Vars()
// 	heroes := AddNewHero()
// 	json.NewEncoder(response).Encode(heroes)
// }
