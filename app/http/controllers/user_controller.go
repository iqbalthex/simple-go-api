package controllers

import (
  "fmt"

  "gate_go/app/facades"
  "gate_go/app/models"
  repo "gate_go/app/repositories"
  "gate_go/app/resources"

  "github.com/gofiber/fiber/v2"
)

type (
  UserController struct {
    userRepo *repo.UserRepository
  }
  UserJson struct{
    Id    uint          `serialize:"hide:true"`
    Name  string
    Posts []models.Post `gorm:"foreignKey:UserId"`
  }
)

func NewUserController(userRepo *repo.UserRepository) *UserController {
  ctrl := &UserController{
    userRepo: userRepo,
  }

  return ctrl
}

func (ctrl *UserController) Index(c *fiber.Ctx) error {
  var users []models.User

  facades.DB().Table("users").Preload("Posts").Find(&users)

  return c.JSON(&fiber.Map{
    "users": resources.NewUserResources(&users, "posts"),
  })
}

func (ctrl *UserController) Store(c *fiber.Ctx) error {
  user := models.User{
    Name: c.FormValue("name"),
    Email: c.FormValue("email"),
  }

  messages := facades.Validate(user)

  if len(messages) > 0 {
    return c.Status(400).JSON(&fiber.Map{
      "errors": messages,
    })
  }

  _, err := ctrl.userRepo.Create(
    user.Name,
    user.Email,
  )

  if err != nil {
    fmt.Println(err)

    return c.Status(500).JSON(&fiber.Map{
      "error": "Internal server error",
    })
  }

  return c.Status(201).JSON(&fiber.Map{
    "message": "User created",
  })
}

func (ctrl *UserController) Show(c *fiber.Ctx) error {
  var user models.User

  id, _ := c.ParamsInt("id")

  facades.DB().Table("users").Preload("Posts").First(&user, id)

  if user.Id < 1 {
    return c.Status(404).JSON(&fiber.Map{
      "error": "User not found",
    })
  }

  return c.JSON(&fiber.Map{
    "user": resources.NewUserResource(&user, "posts"),
  })
}

func (ctrl *UserController) Update(c *fiber.Ctx) error {
  id, _ := c.ParamsInt("id")

  _, err := ctrl.userRepo.Update(
    id,
    c.FormValue("name"),
    c.FormValue("email"),
  )

  if err != nil {
    fmt.Println(err)

    return c.Status(500).JSON(&fiber.Map{
      "error": "Internal server error",
    })
  }

  return c.Status(200).JSON(&fiber.Map{
    "message": "User updated",
  })
}

func (ctrl *UserController) Destroy(c *fiber.Ctx) error {
  id, _ := c.ParamsInt("id")

  err := ctrl.userRepo.Delete(id)

  if err != nil {
    fmt.Println(err)

    return c.Status(500).JSON(&fiber.Map{
      "error": "Internal server error",
    })
  }

  return c.Status(200).JSON(&fiber.Map{
    "message": "User deleted",
  })
}
