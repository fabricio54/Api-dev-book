package usermodels

import (
	"errors"
	"strings"
	"time"
)

// estrutura da entidade usuário
type User struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedIn time.Time `json:"createdin,omitempty"`
}

// funções para validações dos campos
func (user *User) Prepare() error {

	if err := user.validate(); err != nil {
		return err
	}

	user.formartting()

	return nil
}

// método para verificação de campos
func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if user.Email == "" {
		return errors.New("O Email é obrigatório e não pode estar em branco")
	}

	if user.Nick == "" {
		return errors.New("O Nick é obrigatório e não pode estar em branco")
	}

	if user.Password == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}

	return nil
}

// método para formartar campo
func (user *User) formartting() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Nick = strings.TrimSpace(user.Nick)
}
