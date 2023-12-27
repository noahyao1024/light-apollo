package storage

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error

	endpoint, _ := os.LookupEnv("MYSQL_ENDPOINT")
	user, _ := os.LookupEnv("MYSQL_USER")
	password, _ := os.LookupEnv("MYSQL_PASSWORD")

	if endpoint == "" || user == "" || password == "" {
		panic(fmt.Sprintf("failed to connect database, endpoint: %s, user: %s, password: %s", endpoint, user, password))
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", user, password, endpoint, "ApolloConfigDB", "utf8mb4")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func GetDB() *gorm.DB {
	return db
}
