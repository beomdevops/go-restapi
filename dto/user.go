package dto

import "github.com/beomdevops/go-restapi/models"

type CreateUserRequest struct {
	Name string
}

func NewCreateUserRequest(name string) *CreateUserRequest {
	return &CreateUserRequest{
		Name: name,
	}
}

func (c *CreateUserRequest) ToUserEntity() *models.User {
	return models.NewUser(c.Name)
}
