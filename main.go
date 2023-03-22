package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/memochou1993/go-rollbar-example/logger"
)

func init() {
	logger.InitRollbar()
}

func main() {
	r := echo.New()
	r.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			go logger.Error(err.Error())
			return err
		},
	}))

	logger.Error("Hello")

	r.Logger.Fatal(r.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}
