package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jason-adam/appointment-service/internal/appointment"
	"github.com/jason-adam/appointment-service/internal/config"
	"github.com/jason-adam/appointment-service/internal/routes"
)

func main() {
	localConfig := config.Local()

	dbPool, cancel, err := initDatabase(localConfig.DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	defer cancel()

	apptRepo := appointment.NewRepository(
		appointment.WithConnectionPool(dbPool),
		appointment.WithLocation("America/Los_Angeles"),
	)

	apptSvc := appointment.NewService(apptRepo)

	app := fiber.New()

	// Middleware
	app.Use(requestid.New())
	app.Use(logger.New())

	api := app.Group("/api")

	// Routes
	routes.HealthRouter(api)
	routes.AppointmentRouter(api, apptSvc)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		_ = <-c
		log.Println("Gracefully shutting server down....")
		_ = app.Shutdown()
	}()

	if err := app.Listen(fmt.Sprintf(":%s", localConfig.ServerConfig.Port)); err != nil {
		log.Panic(err)
	}

	log.Println("Server has been shutdown")
}

func initDatabase(conf config.DatabaseConfig) (*pgxpool.Pool, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	poolConfig, err := pgxpool.ParseConfig(conf.URL)
	if err != nil {
		cancel()
		return nil, nil, err
	}

	dbPool, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		cancel()
		return nil, nil, err
	}

	return dbPool, cancel, nil
}
