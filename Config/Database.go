package config

import (
	"fmt"
	"log"

	"os"
	"strconv"

	"github.com/jinzhu/gorm"
)

// DB gorm Variable
var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig Config
func BuildDBConfig() *DBConfig {
	port, err := strconv.Atoi(os.Getenv("Port"))
	if err != nil {
		log.Fatal("Error Parsing Port")
	}

	dbConfig := DBConfig{
		Host:     os.Getenv("Host"),
		Port:     port,
		User:     os.Getenv("User"),
		Password: os.Getenv("Password"),
		DBName:   os.Getenv("DBName"),
	}
	return &dbConfig
}

// DbURL Config
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
