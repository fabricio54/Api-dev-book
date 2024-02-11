package security

import (
	"golang.org/x/crypto/bcrypt"
)

// teremos duas funções nesse pacote: uma para receber uma string e colocar um rash nela e outra para quando o usuário logar saber se a senha está correta

// função 1: recebe uma string e faz um hash nela.
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // bcrypt.DefaultCost retorna uma constante padrão para o custo da operação
}

// função 2: compara uma string com hash
func checkPassword(passwordString, passwordHash string) (error) {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordString))
}