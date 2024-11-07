package main

import (
	"aisale/api"
	"aisale/config"
	"aisale/database"
)

func main() {
	config.LoadENV()

	database.Connect()
	defer database.Disconnect()

	api.RouterStart()
}
