package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdwiastika/the-cows-reog-fiber/internal/api"
	"github.com/mdwiastika/the-cows-reog-fiber/internal/config"
	"github.com/mdwiastika/the-cows-reog-fiber/internal/connection"
	"github.com/mdwiastika/the-cows-reog-fiber/internal/repository"
	"github.com/mdwiastika/the-cows-reog-fiber/internal/service"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	userRepository := repository.NewUser(dbConnection)
	userService := service.NewUser(userRepository)

	api.NewUser(app, userService)

	app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
