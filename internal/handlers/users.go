package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type User struct{
	gorm.Model
	ID			int			`json:"id"`
	Username 	string		`json:"username"`
	password 	string		`json:"password"`
	CreatedAt 	time.Time	`json:"createdAt"`
	UpdatedAt 	time.Time	`json:"updatedAt"`
}

func CreateUser(c *fiber.Ctx) error{
	user := User{
		ID: 1,
		Username: "test",
		password: "test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := c.BodyParser(user); err != nil {
		return err
	}
	return c.SendStatus(200)
}

func GetUser(c *fiber.Ctx) error{
	if err := c.BodyParser(User{}); err != nil{
		return c.SendStatus(200)
	}
	return c.JSON("ok")
}

func UpdateUser(c *fiber.Ctx) error{
	if err := c.BodyParser(User{}); err != nil{
		return c.SendStatus(200)
	}
	return c.SendStatus(200)
}

func DeleteUser(c *fiber.Ctx) error{
	if err := c.BodyParser(User{}); err != nil{
		return c.SendStatus(200)
	}
	return c.SendStatus(200)
}
