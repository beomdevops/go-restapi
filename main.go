package main

import (
	"log"

	"github.com/beomdevops/go-restapi/controller"
	"github.com/beomdevops/go-restapi/database"
	"github.com/beomdevops/go-restapi/repository"
	"github.com/beomdevops/go-restapi/service"
	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	posgresDb := database.NewPostgres()

	err := posgresDb.Connection()

	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(posgresDb.PostgresDB)
	userSvc := service.NewUserService(userRepo)
	userController := controller.NewUserController(userSvc)

	app.Get("/users/:userId", func(c *fiber.Ctx) error {
		return userController.FindUserById(c)
	})

	app.Get("/users/:userName", func(c *fiber.Ctx) error {
		return userController.FindUserById(c)
	})
	app.Post("/users", func(c *fiber.Ctx) error {
		return userController.CreateUser(c)
	})
	log.Fatal(app.Listen(":3000"))

}
