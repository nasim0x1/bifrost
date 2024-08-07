package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host             string
	Port             string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
}

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

var Envs = initConfig()
var DBConfig = DatabaseConfig{
	Host:             getEnv("DB_HOST", "localhost"),
	Port:             getEnv("DB_PORT", "3306"),
	DatabaseUser:     getEnv("DB_USER", "root"),
	DatabasePassword: getEnv("DB_PASSWORD", "mypassword"),
	DatabaseName:     getEnv("DB_NAME", "bifrost"),
}

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   getEnv("PORT", "8080"),
		JWTSecret:              getEnv("JWT_SECRET", "not-so-secret-now-is-it?"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}

func (dc *DatabaseConfig) GetDSN() *mysql.Config {
	return &mysql.Config{
		User:                 dc.DatabaseUser,
		Passwd:               dc.DatabasePassword,
		Addr:                 fmt.Sprintf("%s:%s", dc.Host, dc.Port),
		DBName:               dc.DatabaseName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
}
