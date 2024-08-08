package configs

import (
	"fmt"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host             string
	Port             string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
}

var DBConfig = initDatabaseConfig()

func initDatabaseConfig() DatabaseConfig {
	godotenv.Load()
	return DatabaseConfig{
		Host:             getEnv("DB_HOST", "localhost"),
		Port:             getEnv("DB_PORT", "3306"),
		DatabaseUser:     getEnv("DB_USER", "root"),
		DatabasePassword: getEnv("DB_PASSWORD", "mypassword"),
		DatabaseName:     getEnv("DB_NAME", "bifrost"),
	}
}

func (dc *DatabaseConfig) GetConnectionString() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DBConfig.DatabaseUser, DBConfig.DatabasePassword, DBConfig.DatabaseName)

}
