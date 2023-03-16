package controller

import (
	"encoding/json"
	"fmt"
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

func AddNewHero(response http.ResponseWriter, request *http.Request) {
	heroes := models.Heroes{}
	json.NewDecoder(request.Body).Decode(&heroes)
	errs := []string{}

	if heroes.Name == "" {
		errs = append(errs, fmt.Errorf("Name is required").Error())
	}
	if heroes.Power == "" {
		errs = append(errs, fmt.Errorf("Power is required").Error())
	}
	if heroes.Surname == "" {
		errs = append(errs, fmt.Errorf("Surname is required").Error())
	}
	if heroes.Universe == "" {
		errs = append(errs, fmt.Errorf("Universe is required").Error())
	}

	if len(errs) > 0 {
		fmt.Println("oi")
		response.Header().Add("Content-Type", "application/json")
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(errs)
		return
	}

	newHero := models.AddNewHero(heroes)
	json.NewEncoder(response).Encode(newHero)
}
