package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConnection = ""
	Port             = 0
	DriveDataBase    = ""
)

// inializa as variáveis de ambiente
func Load() {
	// criamos um variável para erro
	var erro error

	// carregamos o arquivo .env
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	DriveDataBase = os.Getenv("DB_DRIVE_CONNECTION")
	// convertendo e pegando o valor da porta na variável de ambiente
	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Port = 9000
	}

	// configurando string de conexão com o banco de dados
	StringConnection = fmt.Sprintf(`%s:%s@/%s?charset=utf8&parseTime=True&loc=Local`, os.Getenv("DB_USUARIO"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
}