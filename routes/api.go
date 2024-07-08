package routes

import (
  // "fmt"
  // "log"

  ctrl "gate_go/app/http/controllers"
  repo "gate_go/app/repositories"

  "github.com/gofiber/fiber/v2"
)

func Api(router fiber.Router) {
  userCtrl := ctrl.NewUserController(
    repo.NewUserRepository(),
  )

  api := router.Group("/api")

  api.Get   ("/users",     userCtrl.Index)
  api.Post  ("/users",     userCtrl.Store)
  api.Get   ("/users/:id", userCtrl.Show)
  api.Put   ("/users/:id", userCtrl.Update)
  api.Delete("/users/:id", userCtrl.Destroy)
}
