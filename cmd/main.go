package main

import (
	"log"

	"github.com/nasim0x1/bifrost/cmd/server"
	"github.com/nasim0x1/bifrost/configs"
	"github.com/nasim0x1/bifrost/db"
)

func main() {
	host := configs.Envs.PublicHost
	port := configs.Envs.Port

	database, err := db.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}
	db.InitStorage(database)

	server := server.NewServer(":8080", host, port, database)
	server.Start(true)
}
