package main

import (
	"EventApp/internal/database"
	"EventApp/internal/env"
	_"github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"log"
)

type application struct {
	port int
	jwtSecret string
	models database.Models
}

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	models := database.NewModels(db)
	app := &application{
		port: env.GetEnvInt("PORT", 8080),
		jwtSecret:env.GetEnvString("JWT_SECRET","supersecret"),
		models: models,	
	}

	if err := app.server(); err != nil {
		log.Fatal(err)
	}

	// You can now use 'app' here
	log.Printf("Application started on port %d", app.port)
}