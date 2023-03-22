package logger

import (
	"log"
	"os"

	"github.com/rollbar/rollbar-go"
)

func InitRollbar() {
	rollbar.SetToken(os.Getenv("ROLLBAR_TOKEN"))
	rollbar.SetEnvironment(os.Getenv("APP_ENV"))
}

func Error(message string) {
	log.Println(message)
	if os.Getenv("ROLLBAR_TOKEN") != "" {
		go rollbar.Error(message)
	}
}
