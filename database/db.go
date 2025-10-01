package database

import (
	"log"
	"os"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	sslMode := os.Getenv("SSL_MODE")
	if sslMode == "" {
		sslMode = "require" // padrão produção
	}

	stringDeConexao := "host=" + os.Getenv("HOST") +
		" user=" + os.Getenv("USER") +
		" password=" + os.Getenv("PASSWORD") +
		" dbname=" + os.Getenv("DBNAME") +
		" port=" + os.Getenv("DBPORT") +
		" sslmode=" + sslMode

	log.Println("STRING DE CONEXÃO:", stringDeConexao) // ✅ Agora dentro da função

	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
	if err != nil {
		log.Panicf("Erro ao conectar com banco de dados: %v", err)
	}

	DB.AutoMigrate(&models.Aluno{})
}
