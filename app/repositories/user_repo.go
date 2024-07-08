package repositories

import (
  "gate_go/app/facades"
  "gate_go/app/models"
)

type UserRepository struct {}

var userRepo *UserRepository

func NewUserRepository() *UserRepository {
  if userRepo == nil {
    userRepo = &UserRepository{}
  }

  return userRepo
}

func (r *UserRepository) Create(name, email string) (models.User, error) {
  user := models.User{
    Name: name,
    Email: email,
    RoleId: 2,
  }

  result := facades.DB().Table("users").Create(&user)

  return user, result.Error
}

func (r *UserRepository) Update(id int, name, email string) (models.User, error) {
  user := models.User{
    Name: name,
    Email: email,
  }

  result := facades.DB().Table("users").Where(id).Updates(&user)

  return user, result.Error
}

func (r *UserRepository) Delete(id int) error {
  result := facades.DB().Table("users").Delete(&models.User{}, id)

  return result.Error
}
