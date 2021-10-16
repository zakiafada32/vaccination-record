package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/zakiafada32/vaccination-record/api"
	userController "github.com/zakiafada32/vaccination-record/api/user"
	"github.com/zakiafada32/vaccination-record/app/config"
	userService "github.com/zakiafada32/vaccination-record/business/user"
	repository "github.com/zakiafada32/vaccination-record/modules"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	db := config.ConnectMySQL()
	repository.Migrate(db)

	userRepository := repository.NewUserRepository(db)
	userService := userService.NewUserService(userRepository)
	userController := userController.NewUserController(userService)

	controller := api.Controller{
		User: userController,
	}

	e := echo.New()

	api.Bootstrap(e, &controller)

	// Start server
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	go func() {
		if err := e.Start(":7000"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
