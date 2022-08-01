package main

type login struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

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
