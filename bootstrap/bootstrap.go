package bootstrap

import "github.com/gofiber/fiber/v2/log"

type Application struct {
	Logger *log.Logger
}

func NewInitializeBootsrap() Application {
	app := Application{}
	return app
}
