package main

import (
	"log"
	"bus_booking/config"
	"bus_booking/server"
	

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("Starting Runers App")

	log.Println("Initializig configuration")
	config := config.InitConfig("buses")

	
	log.Println("Initializig database")
	dbHandler := server.InitDatabase(config)


	log.Println("Initializig HTTP sever")
	httpServer := server.InitHttpServer(config, dbHandler)

	httpServer.Start()
}
