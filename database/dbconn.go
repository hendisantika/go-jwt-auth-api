package database

import (
	"fmt"
	"go-jwt-auth-api/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres" //Enter your password for the DB
	dbname   = "jwt-auth-api"
)

var dsn = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
	host, port, user, password, dbname)

var DB *gorm.DB

func DBconn() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db

	db.AutoMigrate(&models.User{}) // we are going to create a models.go file for the User Model.
}
