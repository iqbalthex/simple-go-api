package models

import (
  "gate_go/app/facades"
)

type Role struct{
  Id   uint64
  Name string
}

func init() {
  facades.DB().AutoMigrate(&Role{})
}
