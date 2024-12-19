package config

import (
	"os"
)

func GetDBConfig() string {
	env := os.Getenv("ENV")
	var fullHost string
	host := os.Getenv("DB_HOST")
	if env == "prod" {
		fullHost = host + ".singapore-postgres.render.com"
	} else {
		fullHost = host
	}
	return "host=" + fullHost +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=" + os.Getenv("SSL_MODE") +
		" TimeZone=Asia/Tokyo"
}
