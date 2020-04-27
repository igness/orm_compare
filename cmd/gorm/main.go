package main

import (
	"log"
	"orm-comparison/internal/pkg/comparison"
	"orm-comparison/internal/pkg/comparison/gormrepo"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal(err)
	}

	r := gormrepo.NewGormRepo(db)
	u := comparison.NewUseCase(r)
	user, err := u.CreateUser(comparison.User{
		Name:   "Kyle",
		Gender: "Male",
		Email:  "kyle.jang@buzzvil.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("user: %+v\n", user)
}
