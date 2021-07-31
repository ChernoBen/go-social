package models

import (
	"errors"
	"strings"
	"time"
)

//estrutura representado tabela de usuarios com campos mapeados em json
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdat,omitempty"`
}

//omitempty = omitir se vazio

//função global que e chama os metodos de verificação e formatação
func (u *User) Prepare() error {
	if err := u.isValid(); err != nil {
		return err
	}
	u.wipeSpace()
	return nil
}

//Metodo que verifica se todos os campos estão preenchidos
func (u *User) isValid() error {
	if u.Name == "" {
		return errors.New("Name should not be empty")
	}
	if u.Nick == "" {
		return errors.New("Nick should not be empty")
	}
	if u.Email == "" {
		return errors.New("Email should not be empty")
	}
	if u.Password == "" {
		return errors.New("Password should not be empty")
	}
	return nil
}

//metodo que remove espacos no inicio a no fim de strings
func (u *User) wipeSpace() {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}
