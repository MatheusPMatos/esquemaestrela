package internal

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConectaComBancodeDados() (*gorm.DB, error) {
	dsn := "host=192.168.1.6 user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Discard,
	})
	if err != nil {
		fmt.Println(err.Error())
		log.Panic("Erro ao conectar com banco de dados")
		return nil, err
	}

	fmt.Println("Migrando modelos...")
	if err = DB.AutoMigrate(
		Nota{},
		Enade{},
		Enem{},
		Curso{},
	); err != nil {
		log.Panic("Erro ao migrador modelos para o banco de dados")
		return nil, err
	}

	return DB, nil

}

func SaveNotas(notas []Nota, DB *gorm.DB) error {
	if err := DB.CreateInBatches(&notas, 1000).Error; err != nil {
		return errors.New("erro ao salvar dados no banco")
	}
	return nil
}
