package models

import (
  "gate_go/app/facades"
)

type Post struct{
  Id     uint64
  UserId uint64
  Title  string
  Slug   string
  Desc   string
}

func init() {
  facades.DB().AutoMigrate(&Post{})
}
