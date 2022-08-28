package repository

import (
	"Week4/config"
	"Week4/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func InitRepository(C config.Config) Repository {
	return Repository{
		DB: InitDatabase(C.D),
	}
}

func InitDatabase(D config.DatabaseConfig) *gorm.DB {
	// Connect to database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", D.DATABASE_NAME, D.DATABASE_PASSWORD, D.DATABASE_HOST, D.DATABASE_PORT, D.DATABASE_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	if err := db.AutoMigrate(model.Food{}, model.Drink{}, model.Staff{}, model.Topping{}); err != nil {
		panic(err)
	}
	log.Print("database sucessfully loaded")

	return db
}
