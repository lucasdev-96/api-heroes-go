package models

type Heroes struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Power    string `json:"power"`
	Universe string `json:"universe"`
}

func AllHeroes() []Heroes {
	slice := []Heroes{}
	slice = append(slice, *createHeroes("Bruce Wayne", "Batman", "Tecnologia", "DC"))
	slice = append(slice, *createHeroes("Clark Kent", "SuperMan", " pode voar, tem força descomunal, visão de raio-x, visão de calor, supersopro, superaudição e invunerabilidade", "DC"))
	slice = append(slice, *createHeroes("Logan", "Wolwerine", "Força Sobre-humana, cura", "Marvel"))
	slice = append(slice, *createHeroes("Bruce Banner", "Hulk", "super força extrema.resistência, durabilidade, imunidade e velocidade sobre-humanas. regeneração.invulnerabilidade.super salto de longa distância.onda de concussão.intelecto genial.", "DC"))
	return slice
}

func createHeroes(name, surname, power, universe string) *Heroes {
	return &Heroes{
		Name:     name,
		Surname:  surname,
		Power:    power,
		Universe: universe,
	}
}
