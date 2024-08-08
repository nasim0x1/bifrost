package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/nasim0x1/bifrost/configs"
)

func NewPostgresStorage() (*sql.DB, error) {
	db, err := sql.Open("postgres", configs.DBConfig.GetConnectionString())
	if err != nil {
		return nil, err
	}
	return db, nil

}
func InitStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("DB: Unable to connect to the database: ", err)
	}
	log.Println("DB: Successfully connected!")
}
