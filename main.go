package main

import (
	"youtube-manager/middlewares"
	"youtube-manager/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middlewares.YoutubeService())

	// routes
	routes.Init(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("error loading .env")
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
