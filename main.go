package main

import (
	"fmt"
	"os"

	"github.com/freonL/restoPOS-restAPI/app"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Type string
	Host string
	Port string
	User string
	Pass string
	Name string
	SSL  string
}

func main() {
	godotenv.Load(".env")
	port := os.Getenv("SRV_PORT")

	app := &app.App{}
	app.Initialize()
	fmt.Println("Running server on :" + port)
	app.Run(":" + port)
}
