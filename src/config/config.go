package config

import (
	"fmt"
	"log"
	"os"

	"Kocannn/Blogging-Platform-API.git/src/model"

	_ "github.com/joho/godotenv/autoload"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func ConnectDB() *gorm.DB{

 dbUser := os.Getenv("DB_USER")
 dbPass := os.Getenv("DB_PASSWORD")
 dbName := os.Getenv("DB_NAME")	

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPass, dbName)
	log.Print(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	
	return db
}

func MigrateDB(db *gorm.DB){
	err := db.AutoMigrate(&model.Posts{})
	if err != nil {
		log.Fatal("Failed to migrate database " ,err)
	}
}