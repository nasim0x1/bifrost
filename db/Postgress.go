package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/nasim0x1/bifrost/configs"
)

func NewPostgresStorage() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", configs.DBConfig.DatabaseUser, configs.DBConfig.DatabasePassword, configs.DBConfig.DatabaseName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil

}
func InitStoragePg(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("DB: Unable to connect to the database: ", err)
	}

	log.Println("DB: Successfully connected!")
}
