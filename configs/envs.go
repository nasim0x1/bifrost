package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	Secret                 string
	JWTExpirationInSeconds int64
}

func (c *Config) GetJwtSecret() []byte {
	return []byte(c.Secret)
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   getEnv("PORT", "8080"),
		Secret:                 getEnv("SECRET_KEY", "not-so-secret-now-is-it?"),
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
