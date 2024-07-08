package resources

import (
  "gate_go/app/models"

  "github.com/gofiber/fiber/v2"
)

func NewUserResources(users *[]models.User, preloads ...string) *[]fiber.Map {
  var userResources []fiber.Map

  for _, user := range *users {
    userResource := fiber.Map{
      "name": user.Name,
      "email": user.Email,
    }

    for _, key := range preloads {
      userResource[key] = user.GetRelation(key)
    }

    userResources = append(userResources, userResource)
  }

  return &userResources
}

func NewUserResource(user *models.User, preloads ...string) *fiber.Map {
  userResource := fiber.Map{
    "name": user.Name,
    "email": user.Email,
  }

  for _, key := range preloads {
    userResource[key] = user.GetRelation(key)
  }

  return &userResource
}
