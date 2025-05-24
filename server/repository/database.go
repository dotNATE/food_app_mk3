package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	dsn := "root:password@tcp(database:3306)/food_app?parseTime=true"

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening DB:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	log.Println("Connected to database successfully!")
}
