package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func Log(level Level, message string) {
	godotenv.Load()

	if level.Number() < 0 && os.Getenv("MODE") != "DEBUG" {
		return
	}

	fmt.Printf(
		"%s [%s] %s\n",
		time.Now().Local().Format("2006/01/02 15:04:05"),
		level.String(),
		message,
	)
}
