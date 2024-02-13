package authentication

// utilizamos um alias para facilitar o uso do pacote jwt
import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// usando pacote externo do jwt (json web token)

// função para criação do token
func CreateTokenJWT(userID uint64) (string, error) {
	// a função MapCalims é utilizada para configurar as permissões que o usuário vai ter ao logar na aplicação
	permissions := jwt.MapClaims{}

	// permissão 1: chave authorized com valor true
	permissions["authorized"] = true

	// passando um tempo para o token expirar: 6 horas no caso da nossa aplicação
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix() // basicamamente estamos pegando a hora nesse instante adicionando mais 6 horas e convertendo para Unix (função que pega a hora desde foi criado o unix e traz o resultado em milesegundos)

	permissions["userId"] = userID // passando o id do usuário que veio nos parâmetros da função

	// gerando um secret e depois assinar
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions) // passando o método de assinatura e as permissões

	// retornando o token com assinatura
	return token.SignedString(config.SecretKey) // secret
}
