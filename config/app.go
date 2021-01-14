package config

import "os"

type App struct {
	Name string
	Env  string
}

func New() *App {
	return &App{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}
}
