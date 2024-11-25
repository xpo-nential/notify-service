package config

import (
	"log"

	"github.com/joho/godotenv"
)

type config struct {
	app *app
}

type app struct {
	token   string
	secret  string
	appName string
}

// method

func (c *config) App() *app {
	return c.app
}

func (a *app) GetAppName() string {
	return a.appName
}

func (a *app) GetToken() string {
	return a.token
}

func (a *app) GetSecret() string {
	return a.secret
}

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
