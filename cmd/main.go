package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/templqte/internal/db"
	"github.com/templqte/internal/handlers"
)

func healthcheck(c *fiber.Ctx) error{
	return c.SendString("OK")
}
func setEndpoints(app *fiber.App){
	app.Get("/api/v1/user", handlers.GetUser)
	app.Post("/api/v1/user", handlers.CreateUser)
	app.Delete("/api/v1/user", handlers.DeleteUser)


}
func setUpDB(){
	var err error
	// gorm.Open("sqlite3", handlers.User)
	db.DBConn, err = gorm.Open("sqlite3", "users.db")
	if err != nil{
		panic("failed to connect")
	}
	fmt.Println("DATABASE is connected")
	db.DBConn.AutoMigrate(&handlers.User{})

}
func main(){
	app := fiber.New()
	app.Get("/healthcheck", healthcheck)
	setUpDB()
	defer db.DBConn.Close()
	setEndpoints(app)
	log.Fatal(app.Listen(":8080"))
}