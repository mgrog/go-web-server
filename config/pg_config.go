package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type PgConfig struct {
	host        string
	port        int
	user        string
	password    string
	dname       string
	sslDisabled bool
}

func PgConfigFromEnv() *PgConfig {
	if os.Getenv("DEBUG_MODE") != "true" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		log.Fatal("port value is invalid")
	}

	cfg := PgConfig{
		host:     os.Getenv("DATABASE_HOST"),
		port:     port,
		user:     os.Getenv("DATABASE_USER"),
		password: os.Getenv("DATABASE_SECRET"),
		dname:    os.Getenv("DATABASE_NAME"),
	}
	return &cfg
}

func (c *PgConfig) SslDisabled() *PgConfig {
	c.sslDisabled = true
	return c
}

func (c *PgConfig) GetDsn() string {
	sslmode := "require"
	if c.sslDisabled {
		sslmode = "disable"
	}
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s", c.user, c.password, c.host, c.port, c.dname, sslmode)
}
