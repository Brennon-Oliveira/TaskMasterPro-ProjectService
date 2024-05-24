package db

import (
	"Brennon-Oliveira/TaskMasterPro-ProjectService/internal/models"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

func InitPostgres() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	db, err := gorm.Open("postgres", dsn)
	if err != nil{
		panic(err)
	}
	
	db.AutoMigrate(&models.Project{}, &models.Member{})

	db.LogMode(true)

	return db

}