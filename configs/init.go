package configs

import (
	"log"

	"github.com/joho/godotenv"
)

type IConfig interface {
	App() IAppConfig
}

type config struct {
	app *app
}

type IAppConfig interface {
	GetToken() string
	GetSecret() string
	GetAppName() string
}

type app struct {
	token   string
	secret  string
	appName string
}

// method

func (c *config) App() IAppConfig {
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
