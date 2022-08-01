package main

type patrimonio struct {
	Id         string `json:"id"`
	Tipo       string `json:"tipo"`
	Modelo     string `json:"modelo"`
	Observacao string `json:"observacao"`
}

type message struct {
	Success bool
	Message string
}

type resposta struct {
	Data    interface{}
	Message interface{}
}
