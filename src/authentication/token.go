package authentication

// utilizamos um alias para facilitar o uso do pacote jwt
import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

// validação do token para entrar em rotas protegidas
func ValidatToken(r *http.Request) error {
	// essa função retorna o token
	tokenString := extractToken(r)

	// agora precisamos verificar se secrety key passada no token e válida. logo apos isso ele retornar uma estrutura do tipo *jwt.Token
	token, err := jwt.Parse(tokenString, returnVerificationKey)

	if err != nil {
		return err
	}

	// agora precisamos validar o token propriamente dito
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	// caso não seja válido ele retorna um error
	return errors.New("token inválido")

}

// para fazer a verificação do token precisamos implementar uma função privada para esse pacote
func extractToken(r *http.Request) string {
	// acessando a chave Authorization do cabeçalho
	token := r.Header.Get("Authorization")

	// o formatao correto e vir duas palavras nesse Authorization: Bearer e o token. então precisamos verificar
	if len((strings.Split(token, " "))) == 2 {
		// retorna a posição 1 da string que seria o token propriamente dito
		return strings.Split(token, " ")[1]
	}

	return ""
}

// construindo uma função para pegar o id do usuário nas permissões
func ExtractUserId(r *http.Request) (uint64, error) {

	tokenString := extractToken(r)

	token, err := jwt.Parse(tokenString, returnVerificationKey)

	if err != nil {
		return 0, err
	}
	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return userId, nil
	}

    return 0, errors.New("token inválido")
}

// função criada para verificar se a secret key está no formato valido
func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	// verifica se o método de assinatura esta na familia de métodos do jwt
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
