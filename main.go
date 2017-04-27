package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qor/admin"
	"github.com/qor/qor"
)

type User struct {
	gorm.Model
	Name string
}

type Product struct {
	gorm.Model
	Name        string
	Description string
}


func main() {
  DB, _ := gorm.Open("sqlite3", "demo.db")
  DB.AutoMigrate(&User{}, &Product{})

  Admin := admin.New(&qor.Config{DB: DB})

  Admin.AddResource(&User{})
  Admin.AddResource(&Product{})

  mux := http.NewServeMux()

  Admin.MountTo("/admin", mux)

  fmt.Println("Listening on: 9000")
  http.ListenAndServe(":9000", mux)
}
