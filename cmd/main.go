package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"go-ecommerce/cmd/api"
	"go-ecommerce/config"
	db2 "go-ecommerce/db"
	"log"
)

func main() {
	db, err := db2.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.User,
		Passwd:               config.Envs.Password,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.Name,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8000", db)
	if err := server.Start(); err != nil {
		log.Fatal("Error starting server ", err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to database")
}
