package main

import (
  "fmt"

  "gate_go/app/facades"
  // "gate_go/app/models"

  // "testing"
)

type (
  Post struct{
    UserId uint
    Title  string
  }

  UserSelect struct{
    ID     uint
    Name   string
    RoleId uint
    Role   models.Role
    Posts  []Post `gorm:"foreignKey:UserId"`
  }
)


func main() {
  var user UserSelect

  facades.DB().
    Table("users").
    Preload("Role").
    Preload("Posts", facades.DB()).
    Find(&user)

  fmt.Println(user)
}
