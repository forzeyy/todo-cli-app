package database

import (
	"log"
	"os"

	"github.com/forzeyy/todo-cli-app/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err.Error())
		os.Exit(2)
	}
	/*
		log.Println("Connected to the SQLite database successfully!")
		db.Logger = logger.Default.LogMode(logger.Info)
		log.Println("Running migrations...")
	*/

	db.AutoMigrate(&model.Task{})

	Database = DbInstance{Db: db}
}
