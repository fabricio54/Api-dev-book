package usermodels

import (
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
