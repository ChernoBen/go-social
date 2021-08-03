package models

import (
	"errors"
	"strings"
)

//representação da tabela Articles que guarda publicações de usuarios
type Articles struct {
	ID         uint64 `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	Content    string `json:"content,omitempty"`
	AuthorID   uint64 `json:"author_id,omitempty"`
	AuthorNick uint64 `json:"author_nick,omitempty"`
	Likes      uint64 `json:"likes"`
	CreatedAt  uint64 `json:"createdat,omitempty"`
}

//metodo que verifica e prepara dados da entrada da requisição
func (a *Articles) Prepare() error {
	if err := a.validate(); err != nil {
		return err
	}
	a.format()
	return nil
}

//metodo que valida campos Title e Content
func (a *Articles) validate() error {
	if a.Title == "" {
		return errors.New("Title can not be blank")
	}
	if a.Content == "" {
		return errors.New("Content can not be blank")
	}
	return nil
}

//Metodo que formata campos title e content
func (a *Articles) format() {
	a.Title = strings.TrimSpace(a.Title)
	a.Content = strings.TrimSpace(a.Content)
}
