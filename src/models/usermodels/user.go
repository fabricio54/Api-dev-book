package usermodels

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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
func (user *User) Prepare(stage string) error {

	if err := user.validate(stage); err != nil {
		return err
	}

	if err := user.formartting(stage); err != nil {
		return err
	}

	return nil
}

// método para verificação de campos
func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if user.Email == "" {
		return errors.New("O Email é obrigatório e não pode estar em branco")
	}

	//validando email passado: estamos validando o formato e não se ele existe
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("O e-mail inserido é inválido")
	}

	if user.Nick == "" {
		return errors.New("O Nick é obrigatório e não pode estar em branco")
	}

	if stage == "cadastro" && user.Password == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}

	return nil
}

// método para formartar campo
func (user *User) formartting(stage string) (error) {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Nick = strings.TrimSpace(user.Nick)

	if stage == "cadastro" {

		passwordHash, err := security.Hash(user.Password)

		if err != nil {
			return err
		}

		user.Password = string(passwordHash)
	}

	return nil
}
