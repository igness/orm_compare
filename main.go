package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	_, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Print(err)
	}
}
