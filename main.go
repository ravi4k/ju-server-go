package main

import (
	"ju-go-server/app/models"
	"ju-go-server/db"
	"ju-go-server/server"
	"log"
	"path"
)

func main() {
	dbConfig, err := db.LoadConfig(path.Base("./config"))
	if err != nil {
		log.Fatal(err.Error())
	}

	database := db.InitDB(dbConfig)
	models.DB = database

	appServer := server.NewServer()
	println("fvdfsfdbg")
	err = appServer.Run("3000")
	if err != nil {
		log.Fatal(err.Error())
	}
	println("fvdfsfdbg")
}
