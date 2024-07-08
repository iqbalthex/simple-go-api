package models

import (
  "gate_go/app/facades"
)

type User struct{
  Id     uint64
  Name   string `validate:"required,min=5"`
  Email  string `validate:"required"`
  Posts  []Post
  RoleId uint64
  Role   Role
}

func init() {
  facades.DB().AutoMigrate(&User{})
}

func (u *User) GetRelation(name string) any {
  switch name {
    case "posts":
      return u.Posts
    default:
      panic("Unknown relation name")
  }
}
