package main

import (
	"log"
	"projectbase/config"
	"projectbase/database"
)

func main() {
	cfg := config.LoadConfig()

	// PostgreSQL
	pg := database.ConnectPostgres(
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDB,
	)

	// MongoDB
	mongoClient := database.ConnectMongo(cfg.MongoURI)

	log.Println("Server running on port:", cfg.AppPort)

	// next step: pass DB to repository & route
	_ = pg
	_ = mongoClient
}