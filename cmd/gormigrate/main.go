package main

import (
	"log"
	"orm-comparison/internal/pkg/comparison/repo"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal(err)
	}

	if db.AutoMigrate(&repo.GormUser{}, &repo.GormFood{}, &repo.GormOrder{}).Error != nil {
		log.Fatal(err)
	}

	log.Print("Migrated successfully")
}
