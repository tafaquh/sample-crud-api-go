package main

import (
	"coba-go/app"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	app.StartApplication()
}
