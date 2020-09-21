//main.go
package main

import (
	"fmt"
	"log"

	"materialsetter3d/config"
	"materialsetter3d/models"
	"materialsetter3d/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var err error

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		log.Fatal("Error loading .env file")
	}

	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Stuff{})
	r := routes.SetupRouter()
	// running
	r.Run()
}
