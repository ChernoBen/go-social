package models

//Struct password representa o formato da requisição de alteração de senha
type Password struct {
	New    string `json:"new"`
	Actual string `json:"actual"`
}
