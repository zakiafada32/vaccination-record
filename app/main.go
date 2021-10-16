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
	doctorController "github.com/zakiafada32/vaccination-record/api/doctor"
	hospitalController "github.com/zakiafada32/vaccination-record/api/hospital"
	userController "github.com/zakiafada32/vaccination-record/api/user"
	vaccineController "github.com/zakiafada32/vaccination-record/api/vaccine"
	"github.com/zakiafada32/vaccination-record/app/config"
	doctorService "github.com/zakiafada32/vaccination-record/business/doctor"
	hospitalService "github.com/zakiafada32/vaccination-record/business/hospital"
	userService "github.com/zakiafada32/vaccination-record/business/user"
	vaccineService "github.com/zakiafada32/vaccination-record/business/vaccine"
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

	vaccineRepository := repository.NewVaccineRepository(db)
	vaccineService := vaccineService.NewVaccineService(vaccineRepository)
	vaccineController := vaccineController.NewVaccineController(vaccineService)

	doctorRepository := repository.NewDoctorRepository(db)
	doctorService := doctorService.NewDoctorService(doctorRepository)
	doctorController := doctorController.NewDoctorController(doctorService)

	hospitalRepository := repository.NewHospitalRepository(db)
	hospitalService := hospitalService.NewHospitalService(hospitalRepository)
	hospitalController := hospitalController.NewHospitalController(hospitalService)

	controller := api.Controller{
		User:     userController,
		Vaccine:  vaccineController,
		Doctor:   doctorController,
		Hospital: hospitalController,
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
