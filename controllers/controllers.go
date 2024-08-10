package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hendisantika/go-jwt-auth-api/database"
	"github.com/hendisantika/go-jwt-auth-api/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14) //GenerateFromPassword returns the bcrypt hash of the password at the given cost i.e. (14 in our case).

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	database.DB.Create(&user) //Adds the data to the DB

	return c.JSON(user)

}
