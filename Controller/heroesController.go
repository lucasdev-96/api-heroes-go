package controller

import (
	models "lucasdev-96/Models"
	"net/http"
)

func GetHeroes(response *http.Response, request *http.Request) {
	models.AllHeroes()
}
