package main

import (
	"log"

	"github.com/nasim0x1/bifrost/cmd/server"
	"github.com/nasim0x1/bifrost/configs"
	"github.com/nasim0x1/bifrost/database"
)

func main() {
	host := configs.Envs.PublicHost
	port := configs.Envs.Port

	db, err := database.NewDatabaseStorage()
	if err != nil {
		log.Fatal(err)
	}
	database.InitStorage(db)

	server := server.NewServer(":8080", host, port, db)
	server.Start(true)
}
