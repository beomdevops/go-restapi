package repository

import (
	"errors"

	"github.com/beomdevops/go-restapi/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(id int) (*models.User, error)
	FindByName(name string) (*models.User, error)
	CreateUser(u *models.User) (*models.User, error)
}

func NewUserRepository(cdb *gorm.DB) UserRepository {
	return &userRepository{db: cdb}
}

type userRepository struct {
	db *gorm.DB
}

func (repo *userRepository) CreateUser(u *models.User) (*models.User, error) {

	result := repo.db.Create(u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil

}

func (repo *userRepository) FindById(id int) (*models.User, error) {

	data := &models.User{}

	//result := repo.db.Clauses(dbresolver.Use("read")).Find(data, "id = ?", id)
	result := repo.db.Find(data, "id = ?", id)
	if result.RowsAffected < 1 {
		return nil, errors.New("not found user")
	}
	return data, nil
}
func (repo *userRepository) FindByName(name string) (*models.User, error) {
	data := &models.User{}
	result := repo.db.Where("name = ?", name).First(data)
	if result.RowsAffected < 1 {
		return nil, errors.New("not found user")
	}
	return data, nil
}
