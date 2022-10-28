package main

import (
	"log"
	"path"

	"ju-go-server/app/models"
	"ju-go-server/db"
	"ju-go-server/server"
)

func main() {
	dbConfig, err := db.LoadConfig(path.Base("./config"))
	if err != nil {
		log.Fatal(err.Error())
	}

	database := db.InitDB(dbConfig)
	models.DB = database

	appServer := server.NewServer()
	err = appServer.Run("3000")
	if err != nil {
		log.Fatal(err.Error())
	}
}
