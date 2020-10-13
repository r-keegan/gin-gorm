package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/r-keegan/gin-gorm/Config"
	"github.com/r-keegan/gin-gorm/Models"
	"github.com/r-keegan/gin-gorm/Routes"
)

var err error

func main() {
	Config.DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()

	r.Run()
}
