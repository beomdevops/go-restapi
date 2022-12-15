package controller

import (
	"strconv"

	"github.com/beomdevops/go-restapi/dto"
	"github.com/beomdevops/go-restapi/service"
	fiber "github.com/gofiber/fiber/v2"
)

type UserController struct {
	userSvc *service.UserService
}

func NewUserController(di_usv *service.UserService) *UserController {
	return &UserController{
		userSvc: di_usv,
	}
}

func (userController *UserController) FindUserById(ctx *fiber.Ctx) error {
	parma := ctx.Params("userId")
	id, err := strconv.Atoi(parma)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": err, "data": err})
	}
	find_user, err := userController.userSvc.GetUserById(id)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": err, "data": nil})
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "success", "data": find_user})

}

func (userController *UserController) FindUserByName(ctx *fiber.Ctx) error {
	name := ctx.Params("userName")
	find_user, err := userController.userSvc.GetUserByName(name)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": err, "data": nil})
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "success", "data": find_user})
}

func (userController *UserController) CreateUser(ctx *fiber.Ctx) error {
	p := new(dto.CreateUserRequest)
	err := ctx.BodyParser(p)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": err, "data": err})
	}

	user, err := userController.userSvc.SignUser(p)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "not create", "data": nil})
	}
	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "success", "data": user})

}
