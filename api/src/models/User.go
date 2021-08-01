package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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

//função global que e chama os metodos de verificação e formatação/ parametro step para definir se é uma estapa de criação ou edição de usuario
func (u *User) Prepare(step string) error {
	if err := u.isValid(step); err != nil {
		return err
	}
	if err := u.format(step); err != nil {
		return err
	}
	return nil
}

//Metodo que verifica se todos os campos estão preenchidos
func (u *User) isValid(step string) error {
	if u.Name == "" {
		return errors.New("Name should not be empty")
	}
	if u.Nick == "" {
		return errors.New("Nick should not be empty")
	}
	if u.Email == "" {
		return errors.New("Email should not be empty")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("Invalid email")
	}
	if step == "register" && u.Password == "" {
		return errors.New("Password should not be empty")
	}
	return nil
}

//metodo que remove espacos no inicio a no fim de strings e isere hash no campo password
func (u *User) format(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
	if step == "register" {
		hashPass, err := security.Hash(u.Password)
		if err != nil {
			return err
		}
		u.Password = string(hashPass)
	}
	return nil
}
