package main

import (
	"log"

	"github.com/beomdevops/go-restapi/controller"
	"github.com/beomdevops/go-restapi/database"
	"github.com/beomdevops/go-restapi/models"
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
	posgresDb.PostgresDB.AutoMigrate(&models.User{})
	userRepo := repository.NewUserRepository(posgresDb.PostgresDB)
	userSvc := service.NewUserService(userRepo)
	userController := controller.NewUserController(userSvc)

	app.Get("/users/:userId", func(c *fiber.Ctx) error {
		return userController.FindUserById(c)
	})
	app.Get("users", func(c *fiber.Ctx) error {
		data, err := userRepo.FindUsers()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": err, "data": nil})
		}
		return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "sucess", "data": data})
	})
	app.Get("/users/:userName", func(c *fiber.Ctx) error {
		return userController.FindUserByName(c)
	})
	app.Post("/users", func(c *fiber.Ctx) error {
		return userController.CreateUser(c)
	})

	app.Get("/jwt", func(c *fiber.Ctx) error {
		jwt := service.GenJwt()
		return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "sucess", "jwt": jwt})
	})

	app.Get("/jwk", func(c *fiber.Ctx) error {
		n, e := service.GenJwk()
		return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "sucess", "n": n, "e": e})
	})
	log.Fatal(app.Listen(":3000"))

}
