package main

import (
	"authentication_microservice/database"
	"authentication_microservice/handler"
	"authentication_microservice/server"
	"log"
	"net/http"
	"os"
)

var (
	port         = "8080"
	databaseFile = "./sqliteDB/authenticator.db"
)

func init() {
	if env := os.Getenv("PORT"); env != "" {
		port = env
	}
	if env := os.Getenv("DATABASE_FILE_NAME"); env != "" {
		databaseFile = env
	}
}

func main() {
	db, err := database.InitDB(databaseFile)
	if err != nil {
		panic(err)
	}
	server := server.Server{
		Router: http.NewServeMux(),
	}
	h := handler.Handler{
		DB: db,
	}
	server.InitRoute(&h)
	log.Fatal(http.ListenAndServe(`:`+port, server.Router))

}
