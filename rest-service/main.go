package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("SP// Backend Developer Test - RESTful Service")

	log.Formatter = &logrus.JSONFormatter{}
	e := echo.New()
	// middleware to log incoming requests
	e.Use(middleware.Logger())
	e.GET("/people", getAllPeople)
	e.GET("/people/:id", getPersonByID)

	go func() {
		err := e.Start(":8080")
		if err != nil {
			log.Fatal(err)
		}
	}()

	// graceful server shutdown
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Kill)
	signal.Notify(ch, os.Interrupt)

	<-ch
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()
	e.Shutdown(ctx)

}
