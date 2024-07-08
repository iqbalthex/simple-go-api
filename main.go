package main

import (
  // "fmt"
  "log"

  "gate_go/routes"

  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
  app := fiber.New()

  app.Use(logger.New(logger.Config{
    Format: "${status} - ${method} ${path} (${latency})\n",
  }))

  routes.Api(app)

  log.Fatal(app.Listen(":3000"))
}
