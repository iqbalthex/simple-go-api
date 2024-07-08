package facades

import (
  "fmt"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

var (
  db  *gorm.DB
  err error
)

func init() {
  if db == nil {
    dsn := "root@tcp(127.0.0.1:3306)/gate_go?charset=utf8mb4&parseTime=True&loc=Local"

    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
      PrepareStmt: true,
    })
  }

  if err != nil {
    fmt.Println(err)
    panic(err)
  }
}

func DB() *gorm.DB {
  return db
}
