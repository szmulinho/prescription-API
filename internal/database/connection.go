package database

import (
	"github.com/szmulinho/prescription/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {

	connection, err := gorm.Open(postgres.Open("host=localhost user=postgres password=L96a1prosniper dbname=prescription port=5433 sslmode=disable TimeZone=Europe/Warsaw"), &gorm.Config{})

	if err != nil {
		panic("can't connect with database")
	}

	DB = connection

	connection.AutoMigrate(&model.User{})
	connection.AutoMigrate(&model.CreatePrescInput{})

	return connection
}
