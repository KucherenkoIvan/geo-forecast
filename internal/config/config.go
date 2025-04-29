package config

import (
	"os"
	"strconv"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	// PostgreSQL
	PG_HOST     string
	PG_USER     string
	PG_DB       string
	PG_PASSWORD string
	PG_PORT     string
	// Auth
	ACCEPT_KEYS []string
	// Http
	RESTART_ATTEMPTS int
	LISTEN_PORT      int
}

func Load() {
	val, exist := os.LookupEnv("PG_HOST")
	if exist {
		Values.PG_HOST = val
	} else {
		Values.PG_HOST = "localhost"
	}

	val, exist = os.LookupEnv("PG_USER")
	if exist {
		Values.PG_USER = val
	} else {
		Values.PG_USER = "postgres"
	}

	val, exist = os.LookupEnv("PG_DB")
	if exist {
		Values.PG_DB = val
	} else {
		Values.PG_DB = "postgres"
	}

	val, exist = os.LookupEnv("PG_PASSWORD")
	if exist {
		Values.PG_PASSWORD = val
	} else {
		Values.PG_PASSWORD = "123456"
	}

	val, exist = os.LookupEnv("PG_PORT")
	if exist {
		Values.PG_PORT = val
	} else {
		Values.PG_PORT = "5432"
	}

	val, exist = os.LookupEnv("ACCEPT_KEYS")
	if exist {
		Values.ACCEPT_KEYS = strings.Split(val, ",")
	} else {
		Values.ACCEPT_KEYS = []string{"changeme"}
	}

	val, exist = os.LookupEnv("RESTART_ATTEMPTS")
	if exist {
		val_int, error := strconv.Atoi(val)
		if error == nil {
			Values.RESTART_ATTEMPTS = val_int
		} else {
			Values.RESTART_ATTEMPTS = 25
		}
	} else {
		Values.RESTART_ATTEMPTS = 25
	}

	val, exist = os.LookupEnv("LISTEN_PORT")
	if exist {
		val_int, error := strconv.Atoi(val)
		if error == nil {
			Values.LISTEN_PORT = val_int
		} else {
			Values.LISTEN_PORT = 8080
		}
	} else {
		Values.LISTEN_PORT = 8080
	}

}

// FIXME: Пока что хз как сделать ее readonly,
// поэтому за изменение переменной бьем палкой по хребту
var Values config
